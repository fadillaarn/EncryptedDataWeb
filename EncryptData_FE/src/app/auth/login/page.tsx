"use client";
import React, { useState } from "react";
import { Toast } from "../../dashboard/components/Toast";
import { toast } from "react-toastify";
import Link from "next/link";
import "../styles.css";

export default function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleLogin = (event: { preventDefault: () => void }) => {
    event.preventDefault();
    const url = "http://localhost:8080/api/user/login";
    const formData = new FormData();
    formData.append("email", email);
    formData.append("password", password);

    fetch(url, {
      method: "POST",
      body: formData,
    })
      .then((response) => {
        if (response.ok) {
          return response.json();
        } else {
          toast.error("Login Failed: " + response.statusText);
        }
      })
      .then((data) => {
        if (data) {
          toast.success("Login Success");
          sessionStorage.setItem("token", data.data.token);
          setTimeout(() => {
            window.location.href = "/dashboard";
          }, 2000);
        }
      });
  };
  return (
    <div>
      <Toast />
      <div className="bg-primary">
        <div id="layoutAuthentication">
          <div id="layoutAuthentication_content">
            <main>
              <div className="container">
                <div className="row justify-content-center">
                  <div className="col-lg-5">
                    <div className="card shadow-lg border-0 rounded-lg mt-5">
                      <div className="card-header">
                        <h3 className="text-center font-weight-light my-4">
                          Login
                        </h3>
                      </div>
                      <div className="card-body">
                        <form onSubmit={handleLogin}>
                          <div className="form-floating mb-3">
                            <input
                              className="form-control"
                              type="text"
                              placeholder="Email"
                              required
                              autoComplete="off"
                              value={email}
                              onChange={(e) => setEmail(e.target.value)}
                            />
                            <label>Email</label>
                          </div>
                          <div className="form-floating mb-3">
                            <input
                              className="form-control"
                              type="password"
                              placeholder="Password"
                              required
                              value={password}
                              onChange={(e) => setPassword(e.target.value)}
                            />
                            <label>Password</label>
                          </div>
                          <div className="d-flex align-items-center justify-content-between mt-4 mb-0">
                            <button type="submit" className="btn btn-primary">
                              Login
                            </button>
                          </div>
                        </form>
                      </div>
                      <div className="card-footer text-center py-3">
                        <div className="small">
                          <Link href="/auth/register">
                            Need an account? Sign up!
                          </Link>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </main>
          </div>
        </div>
      </div>
    </div>
  );
}
