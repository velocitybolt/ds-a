import subprocess
import os

def run_openssl(data):
    env = os.environ.copy()
    # the "b" before the string converts it to bytes
    env["password"] = b"asdf"
    proc = subprocess.Popen(
        ["openssl", "enc", "-pbkdf2", "-pass", "env:password"], 
        env=env,
        stdin=subprocess.PIPE,
        stdout=subprocess.PIPE)
    proc.stdin.write(data)
    proc.stdin.flush()
    return proc

procs = []

for _ in range(5):
    data = os.urandom(100)
    proc = run_openssl(data)
    procs.append(proc)

for proc in procs:
    out, _ = proc.communicate()
    print(out)
    
proc = run_openssl(b"apple")
out, _ = proc.communicate()
print(out)