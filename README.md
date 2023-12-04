# portaudio

golang portaudio bindings.

fork of [gordonklaus/portaudio](https://github.com/gordonklaus/portaudio).

## Changes

- PortAudio included.
- Static library.
- Deleted some examples, examples replaced with tests.

## Requirements

- Python 3.11 and above (see [bootstrap.py](./bootstrap.py)).
- CMake 3.1.0 and above (see [CMakeLists.txt](./portaudio/CMakeLists.txt)).

## Windows

- [TDM-GCC](https://jmeubank.github.io/tdm-gcc).
- (not sure) C++ Build Tools and other VS C++ things.

## Other platforms

- gcc.
- **!!! OTHER PLATFORMS NOT TESTED YET !!**.

## Usage

You can include this repository in your project by simply running git clone, or through git submodules.

This project cannot be used via go get, because portaudio has to be built via cmake.

All work is done through the bootstrap.py script. You must run this script before debugging/building your project. For example, in tasks.json if you use VS Code.

bootstrap.py takes the --release argument. In that case, the build will be done without any debug stuff.

Also, by editing this script you can customize CMake flags for portaudio build.

Translated with DeepL.com (free version)

## Updates

You can update portaudio using git submodules. If something changes in portaudio (functions added, removed), you will have to update the bindings manually.

Also cmake options and other such things may change in portaudio. Then you will have to update bootstrap.py as well.
