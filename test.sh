curl -X post 'http://127.0.0.1:6661' \
--data '{
    "req_type": "shell",
    "timeout": 30,
    "params": {
        "spawn": "ssh localhost",
        "interact": [
            {
                "expect": "[Pp]assword:",
                "send": "666666",
                "timeout": 5
            },
            {
                "expect": "[\\$#\\]>]",
                "send": "exit",
                "timeout": 1
            }
        ]
    }
}'
