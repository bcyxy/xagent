# x_agent

```js
// request
{
    "req_type": "shell",
    "timeout": 30,
    "params": {
        "spawn": "ssh localhost",
        "interact": [
            {
                "expect": "[Uu]sername:",
                "send": "yxy",
                "timeout": 10
            },
            {
                "expect": "[Pp]assword:",
                "send": "123456",
                "timeout": 5
            },
            {
                "expect": "[\\$#\\]>]",
                "send": "quit",
                "timeout": 4
            },
            {
                "expect": "[\\$#\\]>]",
                "send": "exit",
                "timeout": 1
            }
        ]
    }
}

// response
{
    "data": [
        "Username:",
        "Password:",
        "<asdf aaa>"
    ]
}
```
