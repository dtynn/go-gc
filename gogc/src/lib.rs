use std::slice::from_raw_parts;

const ALL_ZERO: [u8; 32] = [0; 32];

#[repr(C)]
#[derive(Clone)]
pub struct gogc_Nested {
    pub comm_d: [u8; 32],
}

#[repr(C)]
#[derive(Clone)]
pub struct gogc_PublicReplicaInfo {
    pub comm_r: [u8; 32],
    pub sector_id: u64,
    pub nested: *const *const gogc_Nested,
}

#[repr(C)]
#[derive(Clone)]
pub struct gogc_PublicReplicaInfoCountResponse {
    pub count_slice: *const usize,
    pub count_count: usize,
}

#[no_mangle]
pub unsafe extern "C" fn gogc_verify_post(
    replicas_ptr: *const gogc_PublicReplicaInfo,
    replicas_len: libc::size_t,
) {
    let s = from_raw_parts(replicas_ptr, replicas_len);
    let _ = s.len();
    let base = s[0].sector_id;
    let comm_r = s[0].comm_r;
    let mut m = std::collections::BTreeMap::new();
    s.iter().enumerate().for_each(|(i, r)| {
        assert_ne!(r.sector_id, 0);
        assert!(r.sector_id == (i as u64 + base));
        assert_ne!(r.comm_r, ALL_ZERO);
        assert_eq!(r.comm_r, comm_r);
        m.insert(i, r.clone());
    });

    if s.len() > 0 {
        std::thread::sleep(std::time::Duration::from_millis(50));
    }
}

#[no_mangle]
pub unsafe extern "C" fn gogc_verify_post_count(
    replicas_ptr: *const gogc_PublicReplicaInfo,
    replicas_len: libc::size_t,
) -> *const gogc_PublicReplicaInfoCountResponse {
    let s = from_raw_parts(replicas_ptr, replicas_len);
    let _ = s.len();
    let base = s[0].sector_id;
    let comm_r = s[0].comm_r;
    let counts: Vec<usize> = s
        .iter()
        .enumerate()
        .map(|(i, r)| {
            assert_ne!(r.sector_id, 0);
            assert!(r.sector_id == (i as u64 + base));
            assert_ne!(r.comm_r, ALL_ZERO);
            assert_eq!(r.comm_r, comm_r);
            r.sector_id as usize
        })
        .collect();
    let count_slice = counts.as_ptr();
    let count_count = counts.len();
    std::mem::forget(counts);

    let resp = gogc_PublicReplicaInfoCountResponse {
        count_slice: count_slice,
        count_count: count_count,
    };

    let resp_ptr = (&resp) as *const gogc_PublicReplicaInfoCountResponse;
    std::mem::forget(resp);
    resp_ptr
}

#[no_mangle]
pub unsafe extern "C" fn gogc_count(num: usize) -> usize {
    num
}

#[no_mangle]
pub unsafe extern "C" fn gogc_str_count(_ptr: *const u8, num: usize) -> usize {
    num
}
