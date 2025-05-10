export interface AuthRequest {
    email: string;
    password: string;
}

export interface ResetPasswordPayload {
    password: string;
    accessToken: string,
    refreshToken: string,
}