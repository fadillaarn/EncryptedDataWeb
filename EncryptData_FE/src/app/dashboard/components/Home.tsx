import "../styles/home.css";
import { useEffect, useState } from "react";
import $ from "jquery";
import { toast } from "react-toastify";
import "datatables.net-dt";

export const Home = () => {
  const [allData, setAllData] = useState([]);

  const token = sessionStorage.getItem("token");

  useEffect(() => {
    if (token === null) {
      window.location.href = "/auth/login";
    }
  });

  useEffect(() => {
    const getData = async () => {
      const url = "http://localhost:8080/api/user/getAllMedia";
      const query = await fetch(url, {
        method: "GET",
        headers: {
          Authorization: `Bearer ${token}`,
          "Content-Type": "application/json",
        },
      });
      const response = await query.json();
      setAllData(response);
    };
    getData();
  }, []);

  useEffect(() => {
    const existingTable = $("#filesTable").DataTable();

    if (existingTable) {
      existingTable.destroy();
    }

    const table = $("#filesTable").DataTable({
      data: allData.data,
      columns: [
        {},
        {
          data: "name",
          render: function (data, type, row) {
            if (type === "display" && data.length > 30) {
              return (
                '<span title="' +
                data +
                '">' +
                data.substr(0, 30) +
                "...</span>"
              );
            }
            return data;
          },
        },
        {
          data: "filename",
          render: function (data, type, row) {
            if (type === "display" && data.length > 40) {
              return (
                '<span title="' +
                data +
                '">' +
                data.substr(0, 40) +
                "...</span>"
              );
            }
            return data;
          },
        },
        {
          orderable: false,
          data: null,
          defaultContent: "<button>Download</button>",
        },
      ],
      lengthChange: false,
      columnDefs: [
        {
          targets: 0,
          data: null,
          render: function (data, type, row, meta) {
            return meta.row + 1;
          },
        },
      ],
    });
    table.off("click", "button");

    table.on("click", "button", function (e) {
      let data = table.row(e.target.closest("tr")).data();
      let path = data.path;
      let filename = data.filename;
      var isOk = true;

      console.log(data);

      fetch(path, {
        method: "GET",
        headers: {
          Authorization: "Bearer " + token,
        },
      })
        .then((response) => {
          if (response.ok) {
            toast.success(response.headers.get("Time"));
            return response.blob();
          } else {
            toast.error("Request Failed: " + response.statusText);
            isOk = false;
          }
        })
        .then((blob) => {
          if (isOk) {
            const url = window.URL.createObjectURL(blob);

            const a = document.createElement("a");
            a.href = url;
            a.download = filename;
            a.click();
            window.URL.revokeObjectURL(url);
          }
        });
    });
  }, [allData]);

  return (
    <div className="home">
      <h1>List of All Available Files</h1>
      <table id="filesTable" className="row-border">
        <thead>
          <tr>
            <th>Number</th>
            <th>Owner</th>
            <th>Filename</th>
            <th>Action</th>
          </tr>
        </thead>
        <tbody></tbody>
      </table>
    </div>
  );
};
