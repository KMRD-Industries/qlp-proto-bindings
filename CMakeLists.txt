cmake_minimum_required(VERSION 3.16)

project(qlp-proto LANGUAGES CXX VERSION 0.0.1 DESCRIPTION "c++ protobuf bindings for qlp comms")
set(CMAKE_CXX_EXTENSIONS OFF)
set(CMAKE_CXX_STANDARD 14)
set(CMAKE_CXX_STANDARD_REQUIRED TRUE)

add_library(${PROJECT_NAME} STATIC
    gen/cpp/comm.pb.cc
    gen/cpp/comm.pb.h
)

include(FetchContent)

set(protobuf_INSTALL OFF)
set(protobuf_BUILD_PROTOC_BINARIES OFF)
set(protobuf_BUILD_TESTS OFF)
FetchContent_Declare(protobuf
        GIT_REPOSITORY https://github.com/protocolbuffers/protobuf
        GIT_TAG v3.21.12)
FetchContent_MakeAvailable(protobuf)

target_link_libraries(${PROJECT_NAME} PUBLIC protobuf::libprotobuf)
target_include_directories(${PROJECT_NAME} INTERFACE gen/cpp)

set_target_properties(${PROJECT_NAME} PROPERTIES VERSION ${PROJECT_VERSION})

install(TARGETS ${PROJECT_NAME}
    EXPORT ${PROJECT_NAME}
    LIBRARY DESTINATION ${CMAKE_INSTALL_LIBDIR}
)
