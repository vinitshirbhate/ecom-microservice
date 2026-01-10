package main


type Server struct{

}

// func NewGraphqlServer(accountUrl,catalogUrl,orderUrl) (*Server,error) {
// 	accountClient ,err += accountClient.NewClient(accountUrl)

// 	if err != nil {
// 		return nil,err
// 	}

// 	catalogClient ,err += catalogClient.NewClient(catalogUrl)

// 	if err != nil {
// 		accountClient.close()
// 		return nil,err
// 	}

// 	orderClient ,err += orderClient.NewClient(orderUrl)

// 	if err != nil {
// 		catalogClient.close()
// 		return nil,err
// 	}

// 	return &Server{
// 		accountClient,
// 		catalogClient,
// 		orderClient,
// 	},nil
// }

// func (s *Server) Mutation() MutationResolver{
// 	return &mutationResolver{
// 		server: s,
// 	}
// }

// func (s *Server) Query() QueryResolver{
// 	return &queryResolver{
// 		server: s,
// 	}
// }

// func (s *Server) Account() AccountResolver{
// 	return &accountResolver{
// 		server: s,
// 	}
// }

// func (s *Server) ToExecutableSchema() graphql.ExecutableSchema {
// 	return NewExecutableSchema(Config{
// 		Resolvers: s,
// 	})}