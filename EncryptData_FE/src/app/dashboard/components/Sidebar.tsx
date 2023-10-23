"use client";
import { useState } from "react";
import "../styles/sidebar.css";

const navItems = ["home", "upload", "person", "logout"];

export const Sidebar = ({ onButtonClick }) => {
  const [isOpen, setIsOpen] = useState(false);
  return (
    <nav className={`sidebar ${isOpen ? "open" : ""}`}>
      <div className="sidebar-inner">
        <header className="sidebar-header">
          <button
            type="button"
            className="sidebar-burger"
            onClick={() => setIsOpen(!isOpen)}
          >
            <span className="material-symbols-outlined">
              {isOpen ? "close" : "menu"}
            </span>
          </button>
          <h2>Encryption App</h2>
        </header>
        <nav className="sidebar-menu">
          {navItems.map((item) => (
            <button
              key={item}
              type="button"
              className="sidebar-button"
              onClick={() => onButtonClick(item)}
            >
              <span className="material-symbols-outlined">{item}</span>
              <p>{item === "person" ? "Profile" : item}</p>
            </button>
          ))}
        </nav>
      </div>
    </nav>
  );
};
