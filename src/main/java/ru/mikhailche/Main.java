package ru.mikhailche;

import com.example.grpc.GreetingServiceGrpc;
import com.example.grpc.HelloRequest;
import com.example.grpc.HelloResponse;
import io.grpc.Server;
import io.grpc.ServerBuilder;
import io.grpc.stub.StreamObserver;

import java.io.IOException;

public class Main extends GreetingServiceGrpc.GreetingServiceImplBase {

    public static void main(String[] args) throws IOException, InterruptedException {
        System.out.println("Going to run a server");
        Server server = ServerBuilder.forPort(8080)
                .addService(new Main())
                .build();

        System.out.println("Starting server");
        server.start();

        System.out.println("Awaiting termination");
        server.awaitTermination();
    }

    @Override
    public void greeting(HelloRequest request, StreamObserver<HelloResponse> responseObserver) {
        System.out.println(request);
        responseObserver.onNext(
                HelloResponse.newBuilder()
                        .setGreeting("Hello, " + request.getFirstname())
                        .build()
        );

        responseObserver.onCompleted();
    }
}
