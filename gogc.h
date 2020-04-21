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
  uint64_t count;
} gogc_PublicReplicaInfoResponse;

typedef struct {
  uint8_t comm_r[32];
  uint64_t sector_id;
} gogc_PublicReplicaInfo;

gogc_PublicReplicaInfoResponse gogc_verify_post(const gogc_PublicReplicaInfo *replicas_ptr,
                                                size_t replicas_len);

gogc_PublicReplicaInfoResponse gogc_verify_post_sleep(const gogc_PublicReplicaInfo *replicas_ptr,
                                                      size_t replicas_len);

#endif /* gogc_H */

#ifdef __cplusplus
} /* extern "C" */
#endif

