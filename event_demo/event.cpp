#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <event2/event.h>
#include <event2/listener.h>
#include <event2/bufferevent.h>
#include <event2/buffer.h>

// callback 处理客户端链接
void on_accept(struct evconnlistener *listener, evutil_socket_t fd,
               struct sockaddr *addr, int socklen, void *ctx)
{
    struct event_base *base = (struct event_base *)ctx;
    struct bufferevent *bev = bufferevent_socket_new(base, fd, BEV_OPT_CLOSE_ON_FREE);

    // 设置回调函数，处理客户端发送的数据
    bufferevent_setcb(bev, NULL, NULL, on_event, NULL);
    bufferevent_enable(bev, EV_READ | EV_WRITE);
}

// 定义回调函数，处理客户端发送的数据
void on_event(struct bufferevent *bev, short events, void *ctx)
{
    if (events & BEV_EVENT_EOF)
    {
        printf("Connection closed.\n");
    }
    else if (events & BEV_EVENT_ERROR)
    {
        printf("Some other error.\n");
    }
    else if (events & BEV_EVENT_CONNECTED)
    {
        printf("Connected!\n");
    }

    // 向客户端发送消息
    const char *msg = "Hello, World!\n";
    bufferevent_write(bev, msg, strlen(msg));
}

int main(int argc, char **argv)
{
    struct event_base *base = event_base_new();
    if (!base)
    {
        fprintf(stderr, "Could not initialize libevent!\n");
        return 1;
    }

    struct evconnlistener *listener = evconnlistener_new_bind(base, on_accept, (void *)base,
                                                              LEV_OPT_REUSEABLE | LEV_OPT_CLOSE_ON_FREE, -1,
                                                              (struct sockaddr *)&(struct sockaddr_in){.sin_family = AF_INET, .sin_port = htons(9999)},
                                                              sizeof(struct sockaddr_in));
    if (!listener)
    {
        fprintf(stderr, "Could not create a listener!\n");
        return 1;
    }

    event_base_dispatch(base);

    evconnlistener_free(listener);
    event_base_free(base);
    return 0;
}