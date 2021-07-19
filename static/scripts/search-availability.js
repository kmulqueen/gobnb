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

  let datePick = async function (c) {
    const {
      text = "",
      title = "",
      html = `
      <form id="check-availability-form" action="" method="POST" novalidate class="needs-validation">
        <div class="row">
          <div class="col">
            <div class="row reservation-dates-modal">
              <div class="col">
                <input disabled required class="form-control date-pick-start" type="text" name="start" placeholder="Arrival" />
              </div>
              <div class="col">
                <input disabled required class="form-control date-pick-end" type="text" name="end" placeholder="Departure" />
              </div>
            </div>
          </div>
        </div>
      </form>
    `,
    } = c;

    const { value: formValues } = await Swal.fire({
      title: title,
      html: html,
      focusConfirm: false,
      showCancelButton: true,
      backdrop: false,
      willOpen: () => {
        const elem = document.querySelector(".reservation-dates-modal");
        const rangePicker = new DateRangePicker(elem, {
          showOnFocus: true,
        });
      },
      didOpen: () => {
        document.querySelector(".date-pick-start").removeAttribute("disabled");
        document.querySelector(".date-pick-end").removeAttribute("disabled");
      },
      preConfirm: () => {
        return [
          document.querySelector(".date-pick-start").value,
          document.querySelector(".date-pick-end").value,
        ];
      },
    });

    if (formValues) {
      Swal.fire(JSON.stringify(formValues));
    }
  };

  return {
    toast: toast,
    modal: modal,
    datePick: datePick,
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

//* ============================= Notie Alerts =============================
// type is optional. options are 'success', 'warning', 'error', 'info', 'neutral'
function notify(type, message) {
  notie.alert({
    type: type,
    text: message,
  });
}
