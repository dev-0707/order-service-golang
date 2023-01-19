package users

import (
	"order-service/internal/pkg/models/users"
)

func toUserDto(user users.User) UsersUser {
	userRole := user.Role

	var userRoleDto UsersUserRole = UsersUserRole{
		RoleName: &userRole.RoleName,
	}

	return UsersUser{
		Firstname: &user.Firstname,
		Lastname:  &user.Lastname,
		Role:      &userRoleDto,
		Username:  &user.Username,
	}
}

// func toUser(userDto dto.UsersUser) users.User {
// 	userRoleDto := userDto.Role

// 	var userRole users.UserRole = users.UserRole{RoleName: *userRoleDto.RoleName}

// 	return users.User{
// 		Model:     models.Model{},
// 		Username:  &userDto.Username,
// 		Firstname: &userDto.Firstname,
// 		Lastname:  &userDto.Lastname,
// 		Hash:      "",
// 		Role:      userRole,
// 	}
// }

// func ToProductDTO(product Product) ProductDTO {
// 	return ProductDTO{ID: product.ID, Code: product.Code, Price: product.Price}
// }

// func ToProductDTOs(products []Product) []ProductDTO {
// 	productdtos := make([]ProductDTO, len(products))

// 	for i, itm := range products {
// 		productdtos[i] = ToProductDTO(itm)
// 	}

// 	return productdtos
// }
