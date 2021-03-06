use std::slice::from_raw_parts;

const ALL_ZERO: [u8; 32] = [0u8; 32];

#[repr(C)]
#[derive(Clone)]
pub struct gogc_PublicReplicaInfo {
    pub comm_r: [u8; 32],
    pub sector_id: u64,
}

#[repr(C)]
#[derive(Clone)]
pub struct gogc_PublicReplicaInfoResponse {
    pub count: u64,
}

#[no_mangle]
pub unsafe extern "C" fn gogc_verify_post(
    replicas_ptr: *const gogc_PublicReplicaInfo,
    replicas_len: libc::size_t,
) -> gogc_PublicReplicaInfoResponse {
    let s = from_raw_parts(replicas_ptr, replicas_len);

    let base = s[0].sector_id;
    let mut m = std::collections::BTreeMap::new();
    s.iter().enumerate().for_each(|(i, r)| {
        assert_ne!(r.sector_id, 0);
        assert!(r.sector_id == (i as u64 + base));
        assert_ne!(r.comm_r, ALL_ZERO);
        m.insert(i, r.clone());
    });

    gogc_PublicReplicaInfoResponse {
        count: s.len() as u64,
    }
}

#[no_mangle]
pub unsafe extern "C" fn gogc_verify_post_sleep(
    replicas_ptr: *const gogc_PublicReplicaInfo,
    replicas_len: libc::size_t,
) -> gogc_PublicReplicaInfoResponse {
    let resp = gogc_verify_post(replicas_ptr, replicas_len);
    if resp.count > 0 {
        std::thread::sleep(std::time::Duration::from_millis(50));
    }

    resp
}
