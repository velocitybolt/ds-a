import subprocess
import time

proc = subprocess.Popen(["echo", "Hello from the child!"], stdout=subprocess.PIPE)
out, err = proc.communicate()

print(out.decode("utf-8"))

def hello_fill(input_val):
    print(input_val)

proc2 = subprocess.Popen(["sleep", "0.3"])
while proc2.poll() is None:
    print("Working...")
    time.sleep(0.2)
    hello_fill("LOL")

print("Exit status", proc2.poll())


