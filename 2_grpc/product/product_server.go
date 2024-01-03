package product

import (
	"context"
	"google.golang.org/grpc"
	"grpc_learn/2_grpc/pb"
	"log"
)

// ProductServiceServer 是 Product 服务的具体实现。
type ProductServiceServer struct {
	// 可以在这里添加服务的状态或依赖项
	pb.UnimplementedProductServiceServer
}

// ListProducts 实现了 ProductService 接口中的对应方法。
func (s *ProductServiceServer) ListProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	// 这里应填充实际的产品列表获取逻辑
	return &pb.ListProductsResponse{}, nil // 示例返回空响应
}

// GetProduct 实现了 ProductService 接口中的对应方法。
func (s *ProductServiceServer) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.Product, error) {
	// 这里应填充根据请求参数获取单个产品的逻辑
	log.Println("收到请求:", req.GetId())
	//创建一个产品
	pro := pb.Product{
		Id:       req.GetId(),
		Name:     "手机",
		Price:    1999,
		Image:    "phone.jpg",
		Category: 1,
	}
	return &pro, nil // 示例返回产品
}

// ManageProduct 实现了 ProductService 接口中的对应方法。
func (s *ProductServiceServer) ManageProduct(ctx context.Context, req *pb.ManageProductRequest) (*pb.ManageProductResponse, error) {
	// 这里应填充管理产品（如创建、更新或删除）的逻辑
	return &pb.ManageProductResponse{}, nil // 示例返回空响应
}

// RegisterProductServiceServer 注册 ProductServiceServer 到 gRPC 服务器
func RegisterProductServiceServer(s *grpc.Server, srv *ProductServiceServer) {
	pb.RegisterProductServiceServer(s, srv)
}
