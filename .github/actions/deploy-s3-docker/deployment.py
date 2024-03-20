import os

## This code is just a test to simulate how to create an action with docker container
def run():
    bucket = os.environ['INPUT_BUCKET']
    
    print('bucket test name: ' + bucket)

    toPrint = "hello from python"

    print(f'::set-output name=test_output::{toPrint}')

if __name__ == '__main__':
    run()
