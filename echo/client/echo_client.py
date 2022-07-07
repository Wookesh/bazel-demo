from echo.proto import message_pb2
from echo.proto import echo_pb2
from echo.proto import echo_pb2_grpc
import grpc


def main():
    m = message_pb2.Message(text="pepe")
    req = echo_pb2.EchoRequest(message=m)

    print(f'request: {req}')

    channel = grpc.insecure_channel('localhost:7000')

    stub = echo_pb2_grpc.EchoStub(channel)
    resp = stub.Ping(req)
    print(f'response: {resp}')


if __name__ == '__main__':
    main()
