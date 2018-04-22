vi /etc/yum.repos.d/CentOS-Debuginfo.repo #modify enable=1
yum -y install nss-softokn-debuginfo –nogpgcheck 
yum -y install yum-utils

modify eos/CMakeList.txt 
```
142         # Linux Specific Options Here
143         message( STATUS "Configuring Eos on Linux" )
144         set( CMAKE_CXX_FLAGS "${CMAKE_C_FLAGS} -g -O0 -Wall" )
145         set( FULL_STATIC_BULD "1" )
146         if ( FULL_STATIC_BUILD )
147           set( CMAKE_EXE_LINKER_FLAGS "${CMAKE_EXE_LINKER_FLAGS} -static-libstdc++ -static-libgcc")
148         endif ( FULL_STATIC_BUILD )
149
```
