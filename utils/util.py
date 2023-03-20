def dump(result):
    print(result.choices[0].message.content)


def check_params(**kwargs):
    for k in kwargs:
        if len(kwargs[k])==0:
            print("[!] {} is required".format(k))
            exit(0)
