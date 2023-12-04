import sys
import subprocess
import os
from contextlib import chdir

flags = [
    "-DPA_BUILD_SHARED_LIBS=OFF",
    "-DPA_BUILD_TESTS=OFF",
    "-DPA_BUILD_EXAMPLES=OFF",
    "-DPA_USE_SKELETON=OFF",
    "-DPA_USE_ASIO=OFF",
    "-DPA_USE_DS=ON",
    "-DPA_USE_WMME=ON",
    "-DPA_USE_WASAPI=ON",
    "-DPA_USE_WDMKS_DEVICE_INFO=ON",
    # With flags like this, even if we are using Windows and ALSA is enabled, 
	# it doesn't matter because cmakelists has a check for OS; ALSA, etc. 
	# In short, the flag will still be disabled on systems without ALSA.
    "-DPA_USE_ALSA=ON",
    "-DPA_ALSA_DYNAMIC=OFF",
    "-DPA_USE_OSS=OFF",
    "-DPA_USE_AUDIOIO=OFF",
    "-DPA_USE_PULSEAUDIO=OFF",
]

portaudio_lib = os.path.abspath("./portaudio")
portaudio_build_dir = os.path.abspath("./build")


def main():
    # python bootstrap.py --release
    is_debug = "--release" not in sys.argv
    config = "Release"
    if is_debug == True:
        config = "Debug"
        flags.append("-DPA_ENABLE_DEBUG_OUTPUT=ON")
    else:
        flags.append("-DPA_ENABLE_DEBUG_OUTPUT=OFF")

    # Create build dir.
    if not os.path.exists(portaudio_build_dir):
        os.makedirs(portaudio_build_dir)

    # cd to build dir.
    with chdir(portaudio_build_dir):
        if os.name == "nt":
            # On Windows set MinGW generator, because with MSVC cgo will fail linking.
            subprocess.run(["cmake", portaudio_lib, "-G", "MinGW Makefiles"] + flags)
        else:
            # Use default generator. NOT TESTED.
            subprocess.run(["cmake", portaudio_lib] + flags)
        # There are also questions about the --build and --config flags. But it works on Windows.
        subprocess.run(["cmake", "--build", ".", "--config", config])


if __name__ == "__main__":
    main()
