import subprocess
import os
import sys

ROOT = os.path.abspath(os.path.join(os.path.dirname(__file__), ".."))


def test_imports():
 sys.path.insert(0, ROOT)
 import crontabbuddy
 assert hasattr(crontabbuddy, "__version__")


def test_cli_help():
 result = subprocess.run(
 [sys.executable, "-m", "crontabbuddy", "--help"],
 cwd=ROOT,
 capture_output=True,
 text=True,
 )
 assert result.returncode == 0
 assert "usage" in result.stdout.lower()