package infra

import (
	"bytes"
	c "context"
	"encoding/gob"
	"fmt"
	"time"

	"crud-with-cache/pkg/comment/domain"

	"github.com/redis/go-redis/v9"
)

const (
	cacheTTL      = time.Hour
	bufferChannel = "comment"
)

type redisBuffer struct {
	domain.CommentRepository
	domain.BufferRepository
	redis redis.UniversalClient
}

func NewBuffer(repo domain.CommentRepository, buffer redis.UniversalClient) domain.CommentRepository {
	return &redisBuffer{
		CommentRepository: repo,
		redis:             buffer,
	}
}

func NewSubscriberBuffer(repo domain.BufferRepository, buffer redis.UniversalClient) domain.Subscriber {
	return &redisBuffer{
		BufferRepository: repo,
		redis:            buffer,
	}
}

func (r *redisBuffer) InsertComment(ctx c.Context, comment domain.Comment) error {
	r.buffer(ctx, r.makeKey(comment.FeedID, comment.OwnerID, time.Now()), comment)
	fmt.Println("Comment buffered")
	return nil
}

func (r *redisBuffer) buffer(ctx c.Context, key string, e domain.Comment) {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(e); err != nil {
		//todo log error
		return
	}
	if err := r.redis.Set(ctx, key, buf.Bytes(), cacheTTL); err != nil {
		//todo log error
		return
	}
	if err := r.redis.Publish(ctx, bufferChannel, key).Err(); err != nil {
		//todo log error
		return
	}
}

func (r *redisBuffer) WaitForMessage(ctx c.Context) error {
	sub := r.redis.Subscribe(ctx, bufferChannel)
	ch := sub.Channel()
	array := make([]string, 0, 100)
	timer := time.NewTimer(5 * time.Second)
	fmt.Println("Waiting for message")
	for {
		select {
		case msg := <-ch:
			key := msg.Payload
			array = append(array, key)
			fmt.Println("Message Received: ", key)
			if len(array) >= 100 {
				fmt.Println("timer event : Migration to DB")
				r.migrationToDB(ctx, array)
				array = array[:0]
				timer.Reset(5 * time.Second)
			} else {
				if !timer.Stop() {
					<-timer.C
				}
				timer.Reset(5 * time.Second)
			}
		case <-timer.C:
			if len(array) > 0 {
				fmt.Println("timer event : Migration to DB")
				r.migrationToDB(ctx, array)
				array = array[:0]
			}
			timer.Reset(5 * time.Second)
		case <-ctx.Done():
			return sub.Close()
		}
	}
}

func (r *redisBuffer) migrationToDB(ctx c.Context, messages []string) {
	a := make([]domain.Comment, 0, len(messages))
	for _, key := range messages {
		data, err := r.redis.Get(ctx, key).Bytes()
		if err != nil {
			//todo log error
			continue
		}
		var e domain.Comment
		dec := gob.NewDecoder(bytes.NewReader(data))
		if err := dec.Decode(&e); err != nil {
			//todo log error
			continue
		}
		a = append(a, e)
	}
	if err := r.BufferRepository.InsertCommentWithTransAction(ctx, a); err != nil {
		//todo retry
	}
}

func (r *redisBuffer) makeKey(feedID uint16, userID uint16, bufferedAt time.Time) string {
	return fmt.Sprintf("feed:%d:user:%d:time:%d", feedID, userID, bufferedAt.Unix())
}
