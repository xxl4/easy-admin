# compile for version
make
if [ $? -ne 0 ]; then
    echo "make error"
    exit 1
fi

easy_admin_version=`./easy-admin version`
echo "build version: $easy_admin_version"

# cross_compiles
make -f ./Makefile.cross-compiles

rm -rf ./release/packages
mkdir -p ./release/packages

os_all='linux windows darwin freebsd'
arch_all='386 amd64 arm arm64 mips64 mips64le mips mipsle riscv64'

cd ./release

for os in $os_all; do
    for arch in $arch_all; do
        easy_admin_dir_name="easy-admin_${easy_admin_version}_${os}_${arch}"
        easy_admin_path="./packages/easy-admin_${easy_admin_version}_${os}_${arch}"

        if [ "x${os}" = x"windows" ]; then
            if [ ! -f "./easy-admin_${os}_${arch}.exe" ]; then
                continue
            fi
            if [ ! -f "./easy-admin_${os}_${arch}.exe" ]; then
                continue
            fi
            mkdir ${easy_admin_path}
            mv ./easy-admin_${os}_${arch}.exe ${easy_admin_path}/easy-admin.exe
        else
            if [ ! -f "./easy-admin_${os}_${arch}" ]; then
                continue
            fi
            mkdir ${easy_admin_path}
            mv ./easy-admin_${os}_${arch} ${easy_admin_path}/easy-admin
        fi  
        cp ../LICENSE ${easy_admin_path}
        cp -rf ../config/*.sql ${easy_admin_path}
        cp -rf ../README.md ${easy_admin_path}
        cp -rf ../config/settings.yml ${easy_admin_path}
        cp -rf ../config/settings.full.yml ${easy_admin_path}
        rm -rf ${easy_admin_path}/legacy

        # packages
        cd ./packages
        if [ "x${os}" = x"windows" ]; then
            zip -rq ${easy_admin_dir_name}.zip ${easy_admin_dir_name}
        else
            tar -zcf ${easy_admin_dir_name}.tar.gz ${easy_admin_dir_name}
        fi  
        cd ..
        rm -rf ${easy_admin_path}
    done
done

cd -

echo "Done"
