import openai
from loader.project import *
from loader.loader import *
from manager.select import *

# Act as a web security expert and ready to receive project
SYSTEM_PROMPT_1 = "You are a web security expert and I will send you a project"
# General security assessment
NEED_PROMPT_1 = "Please analyse code above and tell me vulnerabilities in it. Mark every vulnerability with info, warn, medium, high or critical by severity"
# A need prefix to make gpt work better
NEED_PREFIX = "Please analyse code above. "
# Find all taint chains from a given source
NEED_PROMPT_2 = "Can {} become input or parameter of dangerous function calls? Give me the function call chain in format of {}"
# Find all taint chains to a given sink
NEED_PROMPT_3 = "Can remote input in request become input or parameter of {} in a function call chain? Give me the function call chain in format of {}"
# One function call perline format
DEFUALT_TAINT_FORMAT = "one function call per line"
# Editor format
EDITOR_TAINT_FORMAT = "number\n function name\n file name\n line number\n code snippet less than 3 lines\n"
# Semgrep report format
SEMGREP_FORMAT = "semgrep report"
# CodeQL report format
CodeQL_FORMAT = "CodeQL report"

# General security assessment


def need_prompt_1() -> str:
    return NEED_PROMPT_1

# Find all taint chains from a given source


def need_prompt_2(source: str, format=DEFUALT_TAINT_FORMAT) -> str:
    return NEED_PREFIX + NEED_PROMPT_2.format(source, format)

# Find all taint chains to a given sink


def need_prompt_3(sink: str, format=DEFUALT_TAINT_FORMAT) -> str:
    return NEED_PREFIX + NEED_PROMPT_3.format(sink, format)


def match_include(path: str, include: typing.List[str]):
    if len(include) == 0:
        return True
    for v in include:
        if path.startswith(v):
            return True
    return False


def _ask(messages):
    return openai.ChatCompletion.create(
        model="gpt-3.5-turbo",
        messages=messages
    )


def build_message(messages, pro: Project, select: Select, dry=False):
    for path in pro.content:
        if match_include(path, select.include) and path not in select.exclude:
            if dry:
                print(path)
            for k, v in enumerate(pro.content[path]):
                messages.append(
                    {"role": "user", "content": "relative path: {}, part number: {}\n{}".format(path, k, v)})

# add verify=Fasle in openai/api_requestor.py#request_raw L524 to bypass ssl verification


class Manager():
    def set_key(self, api_key: str):
        openai.api_key = api_key

    def set_proxy(self, proxy: str):
        openai.proxy = proxy

    # ask by src, use with load_one
    def ask_src(self, src: typing.List[str]):
        messages = [
            {"role": "system", "content": SYSTEM_PROMPT_1}]
        for chunk in src:
            messages.append({"role": "user", "content": chunk})
        messages.append(
            {"role": "user", "content": NEED_PROMPT_1})

        return _ask(messages)

    # ask by project and select
    def ask(self, pro: Project, select: Select, dry=False):
        messages = [
            {"role": "system", "content": SYSTEM_PROMPT_1}]
        build_message(messages, pro, select, dry)
        messages.append(
            {"role": "user", "content": NEED_PROMPT_1})
        if dry:
            return
        return _ask(messages)

    # ask by project, question and select
    def ask_question(self, pro: Project, select: Select, question: str,  dry=False):
        messages = [
            {"role": "system", "content": SYSTEM_PROMPT_1}]
        build_message(messages, pro, select, dry)
        messages.append(
            {"role": "user", "content": question}
        )
        if dry:
            return
        return _ask(messages)

    # load project by select pack and ask by question
    def execute_task(self, task: Task, dry=False):
        pro = load_project(task.root, task.language)
        return self.ask(pro,  task.select, dry)

    # load project by select pack and ask by question
    def execute_task_question(self, task: Task, question: str, dry=False):
        pro = load_project(task.root, task.language)
        return self.ask_question(pro,  task.select, question, dry)
