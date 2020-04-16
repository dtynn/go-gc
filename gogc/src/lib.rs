use std::slice::from_raw_parts;

#[repr(C)]
#[derive(Clone)]
pub struct gogc_PublicReplicaInfo {
    pub comm_r: [u8; 32],
    pub sector_id: u64,
}

#[no_mangle]
pub unsafe extern "C" fn gogc_verify_post(
    replicas_ptr: *const gogc_PublicReplicaInfo,
    replicas_len: libc::size_t,
) {
    let s = from_raw_parts(replicas_ptr, replicas_len);
    let _ = s.len();
    std::thread::sleep(std::time::Duration::from_millis(120 + s.len() as u64 % 50));
}
