/**
 * @Description: user
 * @Author Lee
 * @Date 2023/12/15 16:29
 **/

package vo

import "time"

type ResSysUser struct {
	Token          string    `json:"token"`
	Isusefree      int64     `json:"isusefree"`
	Type           int64     `json:"type"`
	ExpirationTime time.Time `json:"expiration_time"`
}
