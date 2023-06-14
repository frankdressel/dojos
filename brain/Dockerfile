FROM ubuntu:22.04

RUN apt-get update && apt-get install -y git build-essential cmake curl libncurses5-dev openjdk-18-jdk python3-dev universal-ctags unzip zip
WORKDIR /progs
RUN git clone https://github.com/vim/vim.git && cd vim && ./configure --with-features=huge --enable-multibyte --enable-rubyinterp --enable-python3interp --with-python3-config-dir=$(bash python3-config --configdir) --enable-perlinterp --enable-gui=gtk2 --enable-cscope --prefix=/usr && make VIMRUNTIMEDIR=$(find /usr/share/vim/ -regex '/usr/share/vim/[0-9][0-9]') && make install
RUN git clone --depth 1 https://github.com/fwcd/kotlin-language-server.git && cd kotlin-language-server && sed -i -e 's/javaVersion=.*/javaVersion=18/g' gradle.properties && ./gradlew -Dorg.gradle.java.home=$(find /usr/lib/jvm/ -maxdepth 1 -regex '/usr/lib/jvm/java-[0-9][0-9]-openjdk.*') --no-daemon server:installDist

RUN useradd -ms /bin/bash kotlin
USER kotlin
WORKDIR /home/kotlin
RUN mkdir -p .vim/bundle && git clone https://github.com/VundleVim/Vundle.vim.git .vim/bundle/Vundle.vim
COPY --chown=kotlin .vimrc .vimrc
RUN chmod u+rw .vimrc
RUN vim +PluginInstall +qall
RUN cd .vim/bundle/YouCompleteMe && python3 install.py

USER root
RUN apt-get install -y maven
RUN curl -s "https://get.sdkman.io" | bash
RUN bash -c '. "/root/.sdkman/bin/sdkman-init.sh" && sdk install kotlin'
USER kotlin
WORKDIR /app
