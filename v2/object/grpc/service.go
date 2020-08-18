package object

import (
	refs "github.com/nspcc-dev/neofs-api-go/v2/refs/grpc"
	service "github.com/nspcc-dev/neofs-api-go/v2/service/grpc"
)

// SetAddress sets address of the requested object.
func (m *GetRequest_Body) SetAddress(v *refs.Address) {
	if m != nil {
		m.Address = v
	}
}

// SetRaw sets raw flag of the request.
func (m *GetRequest_Body) SetRaw(v bool) {
	if m != nil {
		m.Raw = v
	}
}

// SetBody sets body of the request.
func (m *GetRequest) SetBody(v *GetRequest_Body) {
	if m != nil {
		m.Body = v
	}
}

// SetMetaHeader sets meta header of the request.
func (m *GetRequest) SetMetaHeader(v *service.RequestMetaHeader) {
	if m != nil {
		m.MetaHeader = v
	}
}

// SetVerifyHeader sets verification header of the request.
func (m *GetRequest) SetVerifyHeader(v *service.RequestVerificationHeader) {
	if m != nil {
		m.VerifyHeader = v
	}
}

// SetObjectId sets identifier of the object.
func (m *GetResponse_Body_Init) SetObjectId(v *refs.ObjectID) {
	if m != nil {
		m.ObjectId = v
	}
}

// SetSignature sets signature of the object identifier.
func (m *GetResponse_Body_Init) SetSignature(v *service.Signature) {
	if m != nil {
		m.Signature = v
	}
}

// SetHeader sets header of the object.
func (m *GetResponse_Body_Init) SetHeader(v *Header) {
	if m != nil {
		m.Header = v
	}
}

// GetChunk returns chunk of the object payload bytes.
func (m *GetResponse_Body_Chunk) GetChunk() []byte {
	if m != nil {
		return m.Chunk
	}

	return nil
}

// SetChunk sets chunk of the object payload bytes.
func (m *GetResponse_Body_Chunk) SetChunk(v []byte) {
	if m != nil {
		m.Chunk = v
	}
}

// SetInit sets initial part of the object.
func (m *GetResponse_Body) SetInit(v *GetResponse_Body_Init) {
	if m != nil {
		m.ObjectPart = &GetResponse_Body_Init_{
			Init: v,
		}
	}
}

// SetChunk sets part of the object payload.
func (m *GetResponse_Body) SetChunk(v *GetResponse_Body_Chunk) {
	if m != nil {
		m.ObjectPart = v
	}
}

// SetBody sets body of the response.
func (m *GetResponse) SetBody(v *GetResponse_Body) {
	if m != nil {
		m.Body = v
	}
}

// SetMetaHeader sets meta header of the response.
func (m *GetResponse) SetMetaHeader(v *service.ResponseMetaHeader) {
	if m != nil {
		m.MetaHeader = v
	}
}

// SetVerifyHeader sets verification header of the response.
func (m *GetResponse) SetVerifyHeader(v *service.ResponseVerificationHeader) {
	if m != nil {
		m.VerifyHeader = v
	}
}

// SetObjectId sets identifier of the object.
func (m *PutRequest_Body_Init) SetObjectId(v *refs.ObjectID) {
	if m != nil {
		m.ObjectId = v
	}
}

// SetSignature sets signature of the object identifier.
func (m *PutRequest_Body_Init) SetSignature(v *service.Signature) {
	if m != nil {
		m.Signature = v
	}
}

// SetHeader sets header of the object.
func (m *PutRequest_Body_Init) SetHeader(v *Header) {
	if m != nil {
		m.Header = v
	}
}

// SetCopiesNumber sets number of the copies to save.
func (m *PutRequest_Body_Init) SetCopiesNumber(v uint32) {
	if m != nil {
		m.CopiesNumber = v
	}
}

// GetChunk returns chunk of the object payload bytes.
func (m *PutRequest_Body_Chunk) GetChunk() []byte {
	if m != nil {
		return m.Chunk
	}

	return nil
}

// SetChunk sets chunk of the object payload bytes.
func (m *PutRequest_Body_Chunk) SetChunk(v []byte) {
	if m != nil {
		m.Chunk = v
	}
}

// SetInit sets initial part of the object.
func (m *PutRequest_Body) SetInit(v *PutRequest_Body_Init) {
	if m != nil {
		m.ObjectPart = &PutRequest_Body_Init_{
			Init: v,
		}
	}
}

// SetChunk sets part of the object payload.
func (m *PutRequest_Body) SetChunk(v *PutRequest_Body_Chunk) {
	if m != nil {
		m.ObjectPart = v
	}
}

// SetBody sets body of the request.
func (m *PutRequest) SetBody(v *PutRequest_Body) {
	if m != nil {
		m.Body = v
	}
}

// SetMetaHeader sets meta header of the request.
func (m *PutRequest) SetMetaHeader(v *service.RequestMetaHeader) {
	if m != nil {
		m.MetaHeader = v
	}
}

// SetVerifyHeader sets verification header of the request.
func (m *PutRequest) SetVerifyHeader(v *service.RequestVerificationHeader) {
	if m != nil {
		m.VerifyHeader = v
	}
}

// SetObjectId sets identifier of the saved object.
func (m *PutResponse_Body) SetObjectId(v *refs.ObjectID) {
	if m != nil {
		m.ObjectId = v
	}
}

// SetBody sets body of the response.
func (m *PutResponse) SetBody(v *PutResponse_Body) {
	if m != nil {
		m.Body = v
	}
}

// SetMetaHeader sets meta header of the response.
func (m *PutResponse) SetMetaHeader(v *service.ResponseMetaHeader) {
	if m != nil {
		m.MetaHeader = v
	}
}

// SetVerifyHeader sets verification header of the response.
func (m *PutResponse) SetVerifyHeader(v *service.ResponseVerificationHeader) {
	if m != nil {
		m.VerifyHeader = v
	}
}

// SetAddress sets address of the object to delete.
func (m *DeleteRequest_Body) SetAddress(v *refs.Address) {
	if m != nil {
		m.Address = v
	}
}

// SetOwnerId sets identifier of the removing object owner.
func (m *DeleteRequest_Body) SetOwnerId(v *refs.OwnerID) {
	if m != nil {
		m.OwnerId = v
	}
}

// SetBody sets body of the request.
func (m *DeleteRequest) SetBody(v *DeleteRequest_Body) {
	if m != nil {
		m.Body = v
	}
}

// SetMetaHeader sets meta header of the request.
func (m *DeleteRequest) SetMetaHeader(v *service.RequestMetaHeader) {
	if m != nil {
		m.MetaHeader = v
	}
}

// SetVerifyHeader sets verification header of the request.
func (m *DeleteRequest) SetVerifyHeader(v *service.RequestVerificationHeader) {
	if m != nil {
		m.VerifyHeader = v
	}
}

// SetBody sets body of the response.
func (m *DeleteResponse) SetBody(v *DeleteResponse_Body) {
	if m != nil {
		m.Body = v
	}
}

// SetMetaHeader sets meta header of the response.
func (m *DeleteResponse) SetMetaHeader(v *service.ResponseMetaHeader) {
	if m != nil {
		m.MetaHeader = v
	}
}

// SetVerifyHeader sets verification header of the response.
func (m *DeleteResponse) SetVerifyHeader(v *service.ResponseVerificationHeader) {
	if m != nil {
		m.VerifyHeader = v
	}
}

// SetAddress sets address of the object with the requested header.
func (m *HeadRequest_Body) SetAddress(v *refs.Address) {
	if m != nil {
		m.Address = v
	}
}

// SetMainOnly sets flag to return the minimal header subset.
func (m *HeadRequest_Body) SetMainOnly(v bool) {
	if m != nil {
		m.MainOnly = v
	}
}

// SetRaw sets raw flag of the request.
func (m *HeadRequest_Body) SetRaw(v bool) {
	if m != nil {
		m.Raw = v
	}
}

// SetBody sets body of the request.
func (m *HeadRequest) SetBody(v *HeadRequest_Body) {
	if m != nil {
		m.Body = v
	}
}

// SetMetaHeader sets meta header of the request.
func (m *HeadRequest) SetMetaHeader(v *service.RequestMetaHeader) {
	if m != nil {
		m.MetaHeader = v
	}
}

// SetVerifyHeader sets verification header of the request.
func (m *HeadRequest) SetVerifyHeader(v *service.RequestVerificationHeader) {
	if m != nil {
		m.VerifyHeader = v
	}
}

// GetShortHeader returns short header of the object.
func (m *HeadResponse_Body_ShortHeader) GetShortHeader() *ShortHeader {
	if m != nil {
		return m.ShortHeader
	}

	return nil
}

// SetShortHeader sets short header of the object.
func (m *HeadResponse_Body_ShortHeader) SetShortHeader(v *ShortHeader) {
	if m != nil {
		m.ShortHeader = v
	}
}

// GetHeader returns object header.
func (m *HeadResponse_Body_Header) GetHeader() *Header {
	if m != nil {
		return m.Header
	}

	return nil
}

// SetHeader sets object header.
func (m *HeadResponse_Body_Header) SetHeader(v *Header) {
	if m != nil {
		m.Header = v
	}
}

// SetHeader sets full header of the object.
func (m *HeadResponse_Body) SetHeader(v *HeadResponse_Body_Header) {
	if m != nil {
		m.Head = v
	}
}

// SetShortHeader sets short header of the object.
func (m *HeadResponse_Body) SetShortHeader(v *HeadResponse_Body_ShortHeader) {
	if m != nil {
		m.Head = v
	}
}

// SetBody sets body of the response.
func (m *HeadResponse) SetBody(v *HeadResponse_Body) {
	if m != nil {
		m.Body = v
	}
}

// SetMetaHeader sets meta header of the response.
func (m *HeadResponse) SetMetaHeader(v *service.ResponseMetaHeader) {
	if m != nil {
		m.MetaHeader = v
	}
}

// SetVerifyHeader sets verification header of the response.
func (m *HeadResponse) SetVerifyHeader(v *service.ResponseVerificationHeader) {
	if m != nil {
		m.VerifyHeader = v
	}
}

// SetMatchType sets match type of the filter.
func (m *SearchRequest_Body_Filter) SetMatchType(v MatchType) {
	if m != nil {
		m.MatchType = v
	}
}

// SetName sets name of the filtering header.
func (m *SearchRequest_Body_Filter) SetName(v string) {
	if m != nil {
		m.Name = v
	}
}

// SetValue sets value of the filtering header.
func (m *SearchRequest_Body_Filter) SetValue(v string) {
	if m != nil {
		m.Value = v
	}
}

// SetVersion sets version of the search query.
func (m *SearchRequest_Body) SetVersion(v uint32) {
	if m != nil {
		m.Version = v
	}
}

// SetFilters sets list of the query filters.
func (m *SearchRequest_Body) SetFilters(v []*SearchRequest_Body_Filter) {
	if m != nil {
		m.Filters = v
	}
}

// SetRaw sets raw flag of the request.
func (m *SearchRequest_Body) SetContainerId(v *refs.ContainerID) {
	if m != nil {
		m.ContainerId = v
	}
}

// SetBody sets body of the request.
func (m *SearchRequest) SetBody(v *SearchRequest_Body) {
	if m != nil {
		m.Body = v
	}
}

// SetMetaHeader sets meta header of the request.
func (m *SearchRequest) SetMetaHeader(v *service.RequestMetaHeader) {
	if m != nil {
		m.MetaHeader = v
	}
}

// SetVerifyHeader sets verification header of the request.
func (m *SearchRequest) SetVerifyHeader(v *service.RequestVerificationHeader) {
	if m != nil {
		m.VerifyHeader = v
	}
}

// SetIdList sets list of the identifiers of the matched objects.
func (m *SearchResponse_Body) SetIdList(v []*refs.ObjectID) {
	if m != nil {
		m.IdList = v
	}
}

// SetBody sets body of the response.
func (m *SearchResponse) SetBody(v *SearchResponse_Body) {
	if m != nil {
		m.Body = v
	}
}

// SetMetaHeader sets meta header of the response.
func (m *SearchResponse) SetMetaHeader(v *service.ResponseMetaHeader) {
	if m != nil {
		m.MetaHeader = v
	}
}

// SetVerifyHeader sets verification header of the response.
func (m *SearchResponse) SetVerifyHeader(v *service.ResponseVerificationHeader) {
	if m != nil {
		m.VerifyHeader = v
	}
}

// SetOffset sets offset of the payload range.
func (m *Range) SetOffset(v uint64) {
	if m != nil {
		m.Offset = v
	}
}

// SetLength sets length of the payload range.
func (m *Range) SetLength(v uint64) {
	if m != nil {
		m.Length = v
	}
}

// SetAddress sets address of the object with the request payload range.
func (m *GetRangeRequest_Body) SetAddress(v *refs.Address) {
	if m != nil {
		m.Address = v
	}
}

// SetRange sets range of the object payload.
func (m *GetRangeRequest_Body) SetRange(v *Range) {
	if m != nil {
		m.Range = v
	}
}

// SetBody sets body of the request.
func (m *GetRangeRequest) SetBody(v *GetRangeRequest_Body) {
	if m != nil {
		m.Body = v
	}
}

// SetMetaHeader sets meta header of the request.
func (m *GetRangeRequest) SetMetaHeader(v *service.RequestMetaHeader) {
	if m != nil {
		m.MetaHeader = v
	}
}

// SetVerifyHeader sets verification header of the request.
func (m *GetRangeRequest) SetVerifyHeader(v *service.RequestVerificationHeader) {
	if m != nil {
		m.VerifyHeader = v
	}
}

// SetChunk sets chunk of the object payload.
func (m *GetRangeResponse_Body) SetChunk(v []byte) {
	if m != nil {
		m.Chunk = v
	}
}

// SetBody sets body of the response.
func (m *GetRangeResponse) SetBody(v *GetRangeResponse_Body) {
	if m != nil {
		m.Body = v
	}
}

// SetMetaHeader sets meta header of the response.
func (m *GetRangeResponse) SetMetaHeader(v *service.ResponseMetaHeader) {
	if m != nil {
		m.MetaHeader = v
	}
}

// SetVerifyHeader sets verification header of the response.
func (m *GetRangeResponse) SetVerifyHeader(v *service.ResponseVerificationHeader) {
	if m != nil {
		m.VerifyHeader = v
	}
}

// SetAddress sets address of the object with the request payload range.
func (m *GetRangeHashRequest_Body) SetAddress(v *refs.Address) {
	if m != nil {
		m.Address = v
	}
}

// SetRanges sets list of the ranges of the object payload.
func (m *GetRangeHashRequest_Body) SetRanges(v []*Range) {
	if m != nil {
		m.Ranges = v
	}
}

// SetSalt sets salt for the object payload ranges.
func (m *GetRangeHashRequest_Body) SetSalt(v []byte) {
	if m != nil {
		m.Salt = v
	}
}

// SetBody sets body of the request.
func (m *GetRangeHashRequest) SetBody(v *GetRangeHashRequest_Body) {
	if m != nil {
		m.Body = v
	}
}

// SetMetaHeader sets meta header of the request.
func (m *GetRangeHashRequest) SetMetaHeader(v *service.RequestMetaHeader) {
	if m != nil {
		m.MetaHeader = v
	}
}

// SetVerifyHeader sets verification header of the request.
func (m *GetRangeHashRequest) SetVerifyHeader(v *service.RequestVerificationHeader) {
	if m != nil {
		m.VerifyHeader = v
	}
}

// SetHashList returns list of the range hashes.
func (m *GetRangeHashResponse_Body) SetHashList(v [][]byte) {
	if m != nil {
		m.HashList = v
	}
}

// SetBody sets body of the response.
func (m *GetRangeHashResponse) SetBody(v *GetRangeHashResponse_Body) {
	if m != nil {
		m.Body = v
	}
}

// SetMetaHeader sets meta header of the response.
func (m *GetRangeHashResponse) SetMetaHeader(v *service.ResponseMetaHeader) {
	if m != nil {
		m.MetaHeader = v
	}
}

// SetVerifyHeader sets verification header of the response.
func (m *GetRangeHashResponse) SetVerifyHeader(v *service.ResponseVerificationHeader) {
	if m != nil {
		m.VerifyHeader = v
	}
}
