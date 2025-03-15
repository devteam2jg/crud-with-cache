# Redis와 Pub,Sub 패턴을 이용한 Look Aside & Write Back 구현
![123](https://github.com/user-attachments/assets/825cc30e-a573-40d8-a658-35b247b29936)
## Scenario
1. 게시글과 게시글에 달린 댓글 기능을 지원함
2. 게시글의 수정 사항은 사용자에게 바로 노출되어야 함
3. 게시글은 읽기 작업이 많고 댓글은 읽기/쓰기 작업이 많은 상황을 가정함
4. 댓글은 사용자에게 실시간 응답을 보장하지 않아도 됨
5. 댓글이 작성된 순서는 보장되어야 함

### Domain definition
+ feed - 게시글
+ comment - 게시글에 달린 댓글

## Feature Implement
![ddd](https://github.com/user-attachments/assets/ab5ce5a9-f898-4b7a-8bf3-dff5821c4479)
+ **Feed 읽기** - 많은 읽기 작업이 필요하므로 **Look Aside** 사용
+ **Feed 쓰기** - 사용자에게 변경 사항이 즉시 노출되어야 하므로 **Write Through** 사용
+ **Comment 읽기** - 많은 읽기 작업이 필요하므로 **Look Aside** 사용
+ **Comment 쓰기** - 실시간 응답을 보장하지 않아도 되고 쓰기 작업이 많기 때문에 **Write Back** 사용

## Details

![123](https://github.com/user-attachments/assets/471eea73-405e-4296-a449-17c76270e757)

**Implement Write Back**
+ Comment는 쓰기 작업이 많기 때문에 buffer를 적용하여 특정 조건을 만족하면 트랜잭션으로 일괄 저장하도록 구현
+ Redis를 buffer와 Message Queue로 사용, 댓글 쓰기 시도 시 buffer에 저장 후 key를 publish   
+ subsciber는 해당 key를 array에 저장함. array가 가득 차거나 특정 시간이 지나면 array에 있는 모든 key를 buffer에서 조회하여 DB에 트랜잭션으로 일괄 저장
+ 새롭게 추가된 Comment들은 기존 Comment cache 만료 시 사용자에게 노출됨

## Infra Architecture
![aaaaa](https://github.com/user-attachments/assets/4ff04102-a4b2-4a19-8d4f-f010ddbd683e)

## Build
```Makefile
# Only UNIX environment
make init-dotenv    # Set env
make up    # Init application
```

