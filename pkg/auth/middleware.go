package auth

//func IsAuthenticated(сfg *config.Config) func(next http.Handler) http.Handler {
//	return func(next http.Handler) http.Handler {
//		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//			c, err := r.Cookie(AccessTokenFieldName)
//			if err != nil {
//				if errors.Is(err, http.ErrNoCookie) {
//					http.Error(w, "Unauthorized", http.StatusUnauthorized)
//					return
//				}
//
//				http.Error(w, "Bad Request", http.StatusBadRequest)
//				return
//			}
//
//			tokenStr := c.Value
//			claims := &jwt.StandardClaims{}
//
//			// Парсим токен и проверяем его валидность
//			token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
//				return "", nil
//			})
//			if err != nil {
//				if errors.Is(err, jwt.ErrSignatureInvalid) {
//					http.Error(w, "Unauthorized", http.StatusUnauthorized)
//					return
//				}
//
//				http.Error(w, "Bad Request", http.StatusBadRequest)
//				return
//			}
//
//			if !token.Valid {
//				http.Error(w, "Unauthorized", http.StatusUnauthorized)
//				return
//			}
//
//			next.ServeHTTP(w, r)
//		})
//	}
//}
