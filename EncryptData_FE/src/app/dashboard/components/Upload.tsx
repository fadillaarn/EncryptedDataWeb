import React, { useEffect, useState, useRef } from "react";
import { toast } from "react-toastify";
import "../styles/upload.css";

export const Upload = () => {
  const token = sessionStorage.getItem("token");
  const url = "http://localhost:8080/api/user/upload/";

  const [selectedFile, setSelectedFile] = useState(null);
  const [selectedFileName, setSelectedFileName] = useState("");
  const fileInputRef = useRef(null);
  const [selectedRadioBtn, setSelectedRadioBtn] = useState("AES");
  const isRadioSelected = (value: string): boolean =>
    selectedRadioBtn === value;
  const handleRadioClick = (e: React.ChangeEvent<HTMLInputElement>): void => {
    setSelectedRadioBtn(e.currentTarget.value);
  };

  useEffect(() => {
    if (token === null) {
      window.location.href = "/auth/login";
    }
  });

  const handleFileChange = (event: any) => {
    const file = event.target.files[0];
    setSelectedFile(file);
    setSelectedFileName(file ? file.name : "");

    const imagePreview = document.getElementById("imagePreview");
    if (file) {
      const reader = new FileReader();
      reader.onload = (e) => {
        imagePreview.src = e.target?.result;
        imagePreview.style.display = "block";
      };
      reader.readAsDataURL(file);
    } else {
      imagePreview.style.display = "none";
    }
  };

  const handleSelectFile = () => {
    if (fileInputRef.current) {
      fileInputRef.current.click();
    }
  };

  const handleSubmit = (event: any) => {
    event.preventDefault();

    if (selectedFile) {
      const formData = new FormData();
      formData.append("media", selectedFile);

      fetch(url + selectedRadioBtn, {
        method: "POST",
        headers: {
          Authorization: "Bearer " + token,
        },
        body: formData,
      })
        .then((response) => {
          if (response.ok) {
            return response.json();
          } else {
            toast.error("Upload Failed: " + response.statusText);
          }
        })
        .then((data) => {
          if (data) {
            const timeTaken = data.data.time;
            toast.success("Your file has been uploaded, " + timeTaken);
            setTimeout(() => {
              window.location.reload();
            }, 3000);
          }
        });
    } else {
      toast.error("File cannot be empty!");
    }
  };

  return (
    <div className="upload">
      <div className="instruction">
        <h1>How to Upload A File</h1>
        <ol>
          <li>This is the upload menu, you can upload your file here.</li>
          <li>Choose your file by clicking the button.</li>
          <li>Choose one of the available encryption method</li>
          <li>Click submit to upload your file</li>
          <li>
            Your file will be uploaded and the filename will be visible to any
            other user in the home menu.
          </li>
          <li>
            In the storage, your file will be encrypted and you will be the only
            one to have the key to decrypt the file.
          </li>
          <li>
            After uploading, you can download your own file in the home menu by
            clicking the download button. But you can only download
            <strong> YOUR OWN FILE </strong>. You can't download other user's
            file.
          </li>
        </ol>
      </div>
      <div className="uploadFile">
        <div className="upload-form">
          <form id="uploadForm" onSubmit={handleSubmit}>
            <label htmlFor="file">Choose Your File Here</label>
            <br />
            <input
              id="file"
              type="file"
              onChange={handleFileChange}
              ref={fileInputRef}
            />
            <div className="chooseFile">
              <h3>{selectedFileName || "Choose File..."}</h3>
              <button
                id="chooseButton"
                type="button"
                onClick={handleSelectFile}
              >
                <span className="material-symbols-outlined">upload</span>
              </button>
            </div>
            <br />
            <label htmlFor="encryption-method">
              Choose One Encryption Method
            </label>
            <div id="encryption-method">
              <input
                type="radio"
                value="AES"
                name="encryption-method"
                checked={isRadioSelected("AES")}
                onChange={handleRadioClick}
              />{" "}
              AES
              <br />
              <input
                type="radio"
                value="RC4"
                name="encryption-method"
                checked={isRadioSelected("RC4")}
                onChange={handleRadioClick}
              />{" "}
              RC4
              <br />
              <input
                type="radio"
                value="DES"
                name="encryption-method"
                checked={isRadioSelected("DES")}
                onChange={handleRadioClick}
              />{" "}
              DES
            </div>
            <br />
            <button id="submitButton" type="submit">
              Submit
            </button>
          </form>
        </div>
        <br />
        <div className="previewFile">
          <h2>File Preview:</h2>
          <div className="image-container">
            <img
              id="imagePreview"
              src=""
              alt="Preview Unavailable"
              style={{ display: "none" }}
            />
          </div>
        </div>
      </div>
    </div>
  );
};
