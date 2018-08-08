/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package types

type MsgHeader struct {
	TxnID     string
	RequestID string
	OPCode    OPCode
}

type OPCode int

const (
	OPInsert OPCode = iota + 1
	OPUpdate
	OPDelete
	OPFind
	OPStartTransaction = 666
	OPCommit           = 667
	OPAbort            = 668
)

type OPINSERT struct {
	MsgHeader                 // 标准报文头
	CollectionName string     // "dbname.collectionname"
	DOCS           []Document // 要插入集合的文档
}

type OPUPDATE struct {
	MsgHeader               // 标准报文头
	CollectionName string   // "dbname.collectionname"
	Update         Document // 指定要执行的更新
	Selector       Document // 文档查询条件
}

type OPDELETE struct {
	MsgHeader               // 标准报文头
	CollectionName string   // "dbname.collectionname"
	Selector       Document // 文档查询条件
}

type OPFIND struct {
	MsgHeader               // 标准报文头
	CollectionName string   // "dbname.collectionname"
	Selector       Document // 文档查询条件
	Start          int64    // start index
	Limit          int64    // limit index
	Sort           string   // sort string
}

type OPSTARTTTRANSATION struct {
	MsgHeader
}

type OPCOMMIT struct {
	MsgHeader
}
type OPABORT struct {
	MsgHeader
}

type ReplyHeader struct {
	MsgHeader
	Processor string
	OK        bool
	Code      bool
	Message   string
}

type OPREPLY struct {
	ReplyHeader              // 标准报文头
	CollectionName string    // "dbname.collectionname"
	Docs           Documents // 文档查询结果
}
