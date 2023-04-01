import os
from loader.project import *
lang_ext = {"go": [".go"], "java": [".java"], "perl": [".pl"], "ruby": [".rb"], "scala": [".scala"], "yaml": [".yaml", ".yml"], "lua": [".lua"], "kotlin": [".kt"], "json": [".json"], "css": [".css"], "rust": [".rs"], "c": [
    ".c"], "c++": [".cpp", ".cc"], "python": [".py"], "javascript": [".js"], "php": [".php"], "html": [".html"], "xml": [".xml"]}

CHUNK_SIZE = 4000


def match_ext(name: str, language: typing.List[str]) -> bool:
    for lang in language:
        if lang in lang_ext:
            for ext in lang_ext[lang]:
                if name.endswith(ext):
                    return True
    return False


def load_one(file: str) -> typing.List[str]:
    with open(file, "r") as f:
        data = f.read()
    return chunk_string(data, CHUNK_SIZE)


def chunk_string(string, n) -> typing.List[str]:
    return [string[i:i+n] for i in range(0, len(string), n)]


def is_test(name: str) -> bool:
    return name.find("test") != -1


def load_project(root: str, language: typing.List[str], skip_test=True) -> Project:
    iter = os.walk(root)
    content = {}
    for e in iter:
        rel_dir = e[0][len(root):].lstrip('/')
        if skip_test and is_test(rel_dir):
            continue
        if len(e[2]):
            for name in e[2]:
                if skip_test and is_test(name):
                    continue
                if match_ext(name, language):
                    rel_path = os.path.join(rel_dir, name)
                    abs_path = os.path.join(e[0], name)
                    content[rel_path] = load_one(abs_path)
    return Project(root, content, language)
