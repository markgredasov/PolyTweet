import $api from '../app/api/api';
import { 
    RegisterRequest, 
    RegisterResponse, 
    LoginRequest, 
    LoginResponse 
} from '../app/api/api';

export class AuthService {
    static async register(data: RegisterRequest): Promise<RegisterResponse> {
        const response = await $api.post<RegisterResponse>('/register', data);
        return response.data;
    }

    static async login(data: LoginRequest): Promise<LoginResponse> {
        const response = await $api.post<LoginResponse>('/login', data);
        return response.data;
    }

    static async dummyLogin(role: 'admin' | 'user'): Promise<LoginResponse> {
        const response = await $api.post<LoginResponse>('/dummyLogin', { role });
        return response.data;
    }

    static logout(): void {
        localStorage.removeItem('token');
    }
}