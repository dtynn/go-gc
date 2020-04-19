/* gogc Header */

#ifdef __cplusplus
extern "C" {
#endif


#ifndef gogc_H
#define gogc_H

/* Generated with cbindgen:0.10.0 */

#include <stdarg.h>
#include <stdbool.h>
#include <stdint.h>
#include <stdlib.h>

typedef struct {
  uint8_t comm_d[32];
} gogc_Nested;

typedef struct {
  uint8_t comm_r[32];
  uint64_t sector_id;
  const gogc_Nested *const *nested;
} gogc_PublicReplicaInfo;

typedef struct {
  const uintptr_t *count_slice;
  uintptr_t count_count;
} gogc_PublicReplicaInfoCountResponse;

uintptr_t gogc_count(uintptr_t num);

uintptr_t gogc_str_count(const uint8_t *_ptr, uintptr_t num);

void gogc_verify_post(const gogc_PublicReplicaInfo *replicas_ptr, size_t replicas_len);

const gogc_PublicReplicaInfoCountResponse *gogc_verify_post_count(const gogc_PublicReplicaInfo *replicas_ptr,
                                                                  size_t replicas_len);

#endif /* gogc_H */

#ifdef __cplusplus
} /* extern "C" */
#endif

