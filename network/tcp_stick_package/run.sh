#!/bin/bash

# 定义 tmux 会话名称
SESSION_NAME="go_project"

# 检查是否已经存在同名 tmux 会话
tmux has-session -t $SESSION_NAME 2>/dev/null
if [ $? != 0 ]; then
  # 创建新的 tmux 会话并命名
  tmux new-session -d -s $SESSION_NAME

  # 在会话中创建分屏
  tmux split-window -h -t $SESSION_NAME

  # 在左侧窗口运行 server
  tmux send-keys -t $SESSION_NAME:0.0 "go run ./server" C-m

  # 在右侧窗口运行 client
  tmux send-keys -t $SESSION_NAME:0.1 "go run ./client" C-m
fi

# 进入 tmux 会话
tmux attach-session -t $SESSION_NAME


# source：https://www.cnblogs.com/jojop/p/14376423.html
# 可以发现输出的结果并没有像客户端发送的次数一样，原先在客户端发送20次的代码在服务端只接收了两次。这看起来像是 TCP 发送的包被粘住了一样，故而产生了所谓“粘包”的问题。
# 这里为代码做下总结，“粘包”问题的缘由可能发生在发送端也可能发生在接收端：

# 由 Nagle 算法造成的发送端的粘包：Nagle 算法是一种改善网络传输效率的算法。简单来说就是当我们提交一段数据给 TCP 发送时，TCP 并不立刻发送此段数据，而是等待一小段时间看看在等待期间是否还有要发送的数据，若有则会一次把这好几段数据发送出去。

# 接收端接受不及时造成的接收端粘包：TCP 会把接收到的数据存在自己的缓冲区中，然后通应用层取数据。当应用层由于某些原因不能及时地把 TCP 的数据取出来，就会造成 TCP 缓冲区中存放了几段数据。