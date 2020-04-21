rust_target_dir=$(cargo metadata --format-version=1 --manifest-path=./gogc/Cargo.toml | jq -r ".target_directory")
cargo build --release --manifest-path=./gogc/Cargo.toml
echo "rust target dir $rust_target_dir"
cp $rust_target_dir/release/libgogc.a ./
find -L $rust_target_dir/ -type f -name "gogc.h" -exec cp -- "{}" . \;
