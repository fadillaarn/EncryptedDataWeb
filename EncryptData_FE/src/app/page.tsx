"use client";
import { useEffect } from "react";
export default function Home() {
  useEffect(() => {
    const token = localStorage.getItem("token");

    if (token !== null) {
      setTimeout(() => {
        window.location.href = "/dashboard";
      }, 0);
    } else {
      setTimeout(() => {
        window.location.href = "/auth/login";
      }, 0);
    }
  }, []);
  return <main></main>;
}
