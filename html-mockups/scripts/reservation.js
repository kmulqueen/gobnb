//* ============================= Prompt/Modal Function Access =============================
/**
 * * Prompt uses sweet alert modals
 */
function Prompt() {
  let toast = function (title, icon = "success", position = "top-end") {
    const Toast = Swal.mixin({
      toast: true,
      position,
      showConfirmButton: false,
      timer: 3000,
      timerProgressBar: true,
      didOpen: (toast) => {
        toast.addEventListener("mouseenter", Swal.stopTimer);
        toast.addEventListener("mouseleave", Swal.resumeTimer);
      },
    });

    Toast.fire({
      title,
      icon,
    });
  };

  let modal = function (
    title,
    text,
    type = "success",
    confirmButtonText = "OK"
  ) {
    Swal.fire({
      title,
      text,
      icon: type,
      confirmButtonText,
    });
  };

  return {
    toast: toast,
    modal: modal,
  };
}

let attention = Prompt();

//* ============================= Form Validation =============================
// Disabling form submissions if there are invalid fields
(function () {
  "use strict";

  // Fetch all the forms we want to apply custom Bootstrap validation styles to
  let forms = document.querySelectorAll(".needs-validation");

  // Loop over them and prevent submission
  Array.prototype.slice.call(forms).forEach(function (form) {
    form.addEventListener(
      "submit",
      function (event) {
        if (!form.checkValidity()) {
          event.preventDefault();
          event.stopPropagation();
        }

        form.classList.add("was-validated");
      },
      false
    );
  });
})();

//* ============================= Date Range Picker =============================
const elem = document.querySelector(".date-range-picker");
const rangepicker = new DateRangePicker(elem, {
  // ...options
});

//* ============================= Notie Alerts =============================
// type is optional. options are 'success', 'warning', 'error', 'info', 'neutral'
function notify(type, message) {
  notie.alert({
    type: type,
    text: message,
  });
}
attention.modal("Heyyyy");
