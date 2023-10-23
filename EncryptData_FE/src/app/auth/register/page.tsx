"use client";
import React, { useState } from "react";
import { toast } from "react-toastify";
import { Toast } from "../../dashboard/components/Toast";
import Link from "next/link";
import "../styles.css";

export default function register() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [email, setEmail] = useState("");
  const [passwordConfirm, setPasswordConfirm] = useState("");
  const [image, setImage] = useState(null);

  const handleRegister = (event: { preventDefault: () => void }) => {
    event.preventDefault();
    if (password !== passwordConfirm) {
      alert("Password and confirm password do not match.");
      return;
    }

    if (image) {
      const url = "http://localhost:8080/api/user";
      const formData = new FormData();
      formData.append("KTP", image);
      formData.append("name", username);
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
            toast.error("Register Failed: " + response.statusText);
          }
        })
        .then((data) => {
          if (data) {
            toast.success("Register success, you can log in now");
            setTimeout(() => {
              window.location.href = "/auth/login";
            }, 3000);
          }
        });
    }
  };

  const handleImageUpload = (e: any) => {
    const file = e.target.files[0];
    setImage(file);

    if (file) {
      const reader = new FileReader();

      reader.readAsDataURL(file);
    }
  };

  return (
    <div className="bg-primary">
      <Toast />
      <div id="layoutAuthentication">
        <div id="layoutAuthentication_content">
          <main>
            <div className="container">
              <div className="row justify-content-center">
                <div className="col-lg-7">
                  <div className="card shadow-lg border-0 rounded-lg mt-5">
                    <div className="card-header">
                      <h3 className="text-center font-weight-light my-4">
                        Create Account
                      </h3>
                    </div>
                    <div className="card-body">
                      <form onSubmit={handleRegister}>
                        <div className="form-floating mb-3">
                          <input
                            className="form-control"
                            type="text"
                            placeholder="Username"
                            required
                            value={username}
                            onChange={(e) => setUsername(e.target.value)}
                          />
                          <label>Username</label>
                        </div>
                        <div className="form-floating mb-3">
                          <input
                            className="form-control"
                            type="email"
                            placeholder="Email"
                            required
                            value={email}
                            onChange={(e) => setEmail(e.target.value)}
                          />
                          <label>Email</label>
                        </div>
                        <div className="row mb-3">
                          <div className="col-md-6">
                            <div className="form-floating mb-3 mb-md-0">
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
                          </div>
                          <div className="col-md-6">
                            <div className="form-floating mb-3 mb-md-0">
                              <input
                                className="form-control"
                                type="password"
                                placeholder="Confirm Password"
                                required
                                value={passwordConfirm}
                                onChange={(e) =>
                                  setPasswordConfirm(e.target.value)
                                }
                              />
                              <label>Confirm Password</label>
                            </div>
                          </div>
                        </div>
                        <div className="form-floating mb-3">
                          ID Card
                          <br />
                          <input
                            type="file"
                            accept="image/*"
                            onChange={handleImageUpload}
                          />
                        </div>
                        <div className="mt-4 mb-0">
                          <div className="d-grid">
                            <button
                              type="submit"
                              className="btn btn-primary btn-block"
                            >
                              Create Account
                            </button>
                          </div>
                        </div>
                      </form>
                    </div>
                    <div className="card-footer text-center py-3">
                      <div className="small">
                        <Link href="/auth/login">
                          Have an account? Go to login
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
  );
}
