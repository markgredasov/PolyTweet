import { create } from "zustand";
import { persist } from "zustand/middleware";
import { AuthService } from "../services/AuthService";

interface AuthState {
    isAuth: boolean;
    userId: string | null;
    email: string | null;
    isLoading: boolean;

    login: (email: string, password: string) => Promise<void>;
    register: (email: string, password: string, role: string) => Promise<void>;
    logout: () => void;
}

const initialState: Pick<AuthState, 'userId' | 'email' | 'isAuth' | 'isLoading'> = {
    userId: null,
    email: null,
    isAuth: false,
    isLoading: false,
};

export const useAuthStore = create<AuthState>()(
    persist(
        (set) => ({
            ...initialState,

            async login(email: string, password: string) {
                try {
                    set({ isLoading: true });
                    const response = await AuthService.login({ email, password });
                    localStorage.setItem('token', response.token);
                    set({ email: email, isAuth: true, isLoading: false });
                } catch (e) {
                    console.log(e);
                    set({ isLoading: false });
                    throw e;
                }
            },

            async register(email: string, password: string, role: string) {
                try {
                    set({ isLoading: true });
                    const response = await AuthService.register({ email, password, role });
                    console.log('Registration response:', response);
                    set({ userId: response.id, email: email, isLoading: false });
                } catch (e) {
                    console.log(e);
                    set({ isLoading: false });
                    throw e;
                }
            },

            logout: () => {
                localStorage.removeItem('token');
                set({ isAuth: false, userId: null, email: null });
            }
        }),
        {
            name: 'auth-storage',
            partialize: (state) => ({
                isAuth: state.isAuth,
                userId: state.userId,
                email: state.email
            }),
        }
    )
);