PKG_ROOT=/tmp
VERSION=$TAG
AXIA_ROOT=$PKG_ROOT/axiago-$VERSION

mkdir -p $AXIA_ROOT

OK=`cp ./build/axiago $AXIA_ROOT`
if [[ $OK -ne 0 ]]; then
  exit $OK;
fi
OK=`cp -r ./build/plugins $AXIA_ROOT`
if [[ $OK -ne 0 ]]; then
  exit $OK;
fi


echo "Build tgz package..."
cd $PKG_ROOT
echo "Version: $VERSION"
tar -czvf "axiago-linux-$ARCH-$VERSION.tar.gz" axiago-$VERSION
aws s3 cp axiago-linux-$ARCH-$VERSION.tar.gz s3://$BUCKET/linux/binaries/ubuntu/$RELEASE/$ARCH/
rm -rf $PKG_ROOT/axiago*
