import { AppState } from "react-native";
import { supabase } from "./client";
import { Session, User } from "@supabase/supabase-js";
import { AuthRequest, ResetPasswordPayload } from "@/types/auth";

/**
 * Interface for authentication services, providing methods for user sign-up, login,
 * logout, and password management.
 */
export interface AuthService {
  /**
   * Registers a new user with the provided email and password.
   */
  signUp({ email, password }: AuthRequest): Promise<Session>;

  /**
   * Authenticates a user with the provided email and password.
   */
  login({ email, password }: AuthRequest): Promise<Session>;

  /**
   * Logs the user out of the current session.
   */
  logout(): Promise<void>;

  /**
   * Sends a password recovery email to the provided email address.
   */
  forgotPassword({ email }: { email: string }): Promise<void>;

  /**
   * Resets the user's password to the provided new password.
   */
  resetPassword(payload: ResetPasswordPayload): Promise<User>;
}

/**
 * Implementation of authentication services using Supabase
 */
export class SupabaseAuth implements AuthService {
  async signUp({ email, password }: { email: string; password: string }): Promise<Session> {
    const { data, error } = await supabase.auth.signUp({
      email,
      password,
    });

    if (error) {
      throw new Error(error.message);
    }

    return data.session!;
  }

  async login({ email, password }: { email: string; password: string }): Promise<Session> {
    const { data, error } = await supabase.auth.signInWithPassword({
      email,
      password,
    });

    if (error) {
      throw new Error(error.message);
    }

    return data.session!;
  }

  async logout(): Promise<void> {
    const { error } = await supabase.auth.signOut();

    if (error) {
      throw new Error(error.message);
    }
  }

  async forgotPassword({ email }: { email: string }): Promise<void> {
    const redirectTo = `${window.location.origin}/login/reset-password`; // Adjust as needed
    const { error } = await supabase.auth.resetPasswordForEmail(email, {
      redirectTo,
    });

    if (error) {
      throw new Error(error.message);
    }
  }

  async resetPassword({
    password,
    accessToken,
    refreshToken,
  }: ResetPasswordPayload): Promise<User> {
    const { error: setSessionError } = await supabase.auth.setSession({
      access_token: accessToken,
      refresh_token: refreshToken,
    });

    if (setSessionError) {
      throw new Error(setSessionError.message);
    }

    const { data, error } = await supabase.auth.updateUser({
      password,
    });

    if (error) {
      throw new Error(error.message);
    }

    return data.user;
  }
}

AppState.addEventListener("change", (state) => {
  if (state === "active") {
    supabase.auth.startAutoRefresh();
  } else {
    supabase.auth.stopAutoRefresh();
  }
});
