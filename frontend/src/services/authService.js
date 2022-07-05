import jwtDecode from "jwt-decode";

class AuthService {
  isUserLoggedIn() {
    return localStorage.getItem("jwt") !== null;
  }

  decodeJwt() {
    return jwtDecode(localStorage.getItem("jwt"));
  }

  setJwt(jwt) {
    localStorage.setItem("jwt", jwt);
  }

  removeJwt() {
    localStorage.removeItem("jwt");
  }

  getJwtField(fieldName) {
    const jwt = this.decodeJwt();
    if (fieldName === "username") return jwt.sub;
    else return jwt[fieldName];
  }
}

export const authService = new AuthService();
