import typing
def common_scan(root:str, language:typing.List[str], include:typing.List[str], exclude=[], key="", proxy="", dry=False):
    import manager
    mgr = manager.Manager()
    if len(key):
        mgr.set_key(key)
    if len(proxy):
        mgr.set_proxy(proxy)
    task = manager.Task(root, language, include, exclude)
    return mgr.execute_task(task, dry)
    
# test dry run
if __name__ == "__main__":
    import sys
    import os

    module_path = os.path.abspath(os.path.join(os.path.dirname(__file__), "../"))
    sys.path.append(module_path)


    def dump(result):
        print(result.choices[0].message.content)

    import manager
    import loader
    import utils
    # output security report
    mgr = manager.Manager()

    mgr.set_key(os.environ.get("OPENAI_API_KEY"))

    mgr.set_proxy("http://127.0.0.1:7890")

    project_root = os.path.join(module_path, "benchmark")

    pro = loader.load_project(project_root, ["python"])

    select = manager.Select(["directory", "include.py"], [
        "directory/exclude.py"])

    # change dry here to send request
    dry = True
    if dry:
        mgr.ask(pro, select, dry)
    else:
        res = mgr.ask(pro, select, dry)
        utils.dump(res)
