import typing
def taint_sink_scan(root:str, language:typing.List[str], sink:str, include=[], exclude=[],  key="", proxy="", dry=False):
    import manager
    import loader
    mgr = manager.Manager()
    if len(key):
        mgr.set_key(key)
    if len(proxy):
        mgr.set_proxy(proxy)
    pro = loader.load_project(root, language)
    select = manager.Select(include, exclude)
    return mgr.ask_question(pro, select,  manager.need_prompt_3(sink), dry)


# test dry run
if __name__ == "__main__":
    import sys
    import os

    module_path = os.path.abspath(
        os.path.join(os.path.dirname(__file__), "../"))
    sys.path.append(module_path)

    import manager
    import loader
    import utils
    # output taint track to sink
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
        mgr.ask_question(pro, select,  manager.need_prompt_3("exif tool"), dry)
    else:
        res = mgr.ask_question(
            pro, select,  manager.need_prompt_3("exif tool"), dry)
        utils.dump(res)
