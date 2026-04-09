export interface RegisterResponse {
    id: string;
    email: string;
    role: string;
    createdAt: string;
}

export interface RegisterRequest {
    email: string;
    password: string;
    role: string;
}

export interface LoginRequest {
    email: string;
    password: string;
}

export interface LoginResponse {
    token: string;
}

export interface ErrorResponse {
    code: string;
    message: string;
}