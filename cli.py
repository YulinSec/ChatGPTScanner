import module
import utils


class ChatGPTScan():
    """
    ChatGPTScan help summary page

    A white box code scan powered by ChatGPT

    Example:

        python chatgptscan.py common_scan --project ./benchmark --language "['python']" --include "['directory']" --proxy http://127.0.0.1:7890

        python chatgptscan.py common_scan --project ./go-sec-code --language "['go']" --include "['controllers/cmdi.go','utils']"  --proxy http://127.0.0.1:8080

        python chatgptscan.py taint_sink_scan --project ./benchmark --language "['python']" --sink "os.system()"  --exclude "['directory/exclude.go']"

    Note:
        --project       path to target project
        --language      languages of the project, decide which file extension will be loaded
        --include       files send to ChatGPT, relative directory or relative filepath, match by prefix 
        --exclude       files not send to ChatGPT, relative directory or relative filepath, match by prefix 
        --sink          decrible your sink, only works in taint_sink_scan
        --key           openai api key, also get from environment variable OPENAI_API_KEY
        --proxy         openai api proxy
        --dry           dry run, not send files to ChatGPT

    """

    def common_scan(self, project: str = "", language: list[str] = [], include: list[str] = [], exclude: list[str] = [],  key="", proxy="", dry=False):
        """
        scan project file and output report
        """

        utils.check_params(project=project, language=language)

        res = module.common_scan(project, language, include,
                                 exclude, key, proxy, dry)
        if res:
            utils.dump(res)

    def taint_sink_scan(self, project: str = "", language: list[str] = [], sink: str = "", include: list[str] = [], exclude: list[str] = [],  key="", proxy="", dry=False):
        """
        scan project and output taint path to sink
        """
        utils.check_params(project=project, language=language, sink=sink)

        res = module.taint_sink_scan(
            project, language, sink, include, exclude, key, proxy, dry)

        if res:
            utils.dump(res)
