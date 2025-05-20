document.addEventListener('DOMContentLoaded', function() {
    if (!localStorage.getItem('token')) {
        window.location.href = '/';
        return;
    }

    const appointmentForm = document.getElementById('appointmentForm');
    if (appointmentForm) {
        appointmentForm.addEventListener('submit', handleCreateAppointment);
    }

    loadAppointments();
});

async function loadAppointments() {
    try {
        const appointments = await api.get('/api/appointments');
        renderAppointments(appointments);
    } catch (error) {
        console.error('Error loading appointments:', error);
    }
}

function renderAppointments(appointments) {
    const container = document.getElementById('appointmentsContainer');
    container.innerHTML = '';

    if (appointments.length === 0) {
        container.innerHTML = '<p>No appointments found.</p>';
        return;
    }

    appointments.forEach(appointment => {
        const card = document.createElement('div');
        card.className = 'appointment-card';
        card.innerHTML = `
            <h3>${appointment.title}</h3>
            <p><strong>Doctor:</strong> ${appointment.doctorName}</p>
            <p><strong>Patient:</strong> ${appointment.patientName}</p>
            <p><strong>Start:</strong> ${new Date(appointment.startTime).toLocaleString()}</p>
            <p><strong>End:</strong> ${new Date(appointment.endTime).toLocaleString()}</p>
            <div class="appointment-actions">
                <button class="btn edit-btn" data-id="${appointment.id}">Edit</button>
                <button class="btn delete-btn" data-id="${appointment.id}">Delete</button>
            </div>
        `;
        container.appendChild(card);
    });

    // Add event listeners to buttons
    document.querySelectorAll('.edit-btn').forEach(btn => {
        btn.addEventListener('click', handleEditAppointment);
    });

    document.querySelectorAll('.delete-btn').forEach(btn => {
        btn.addEventListener('click', handleDeleteAppointment);
    });
}

async function handleCreateAppointment(e) {
    e.preventDefault();

    const title = document.getElementById('appointmentTitle').value;
    const doctorName = document.getElementById('doctorName').value;
    const patientName = document.getElementById('patientName').value;
    const startTime = document.getElementById('startTime').value;
    const endTime = document.getElementById('endTime').value;

    try {
        await api.post('/api/appointments', {
            title,
            doctorName,
            patientName,
            startTime: new Date(startTime).toISOString(),
            endTime: new Date(endTime).toISOString()
        });

        document.getElementById('appointmentForm').reset();
        loadAppointments();
    } catch (error) {
        console.error('Error creating appointment:', error);
        alert('Failed to create appointment');
    }
}

async function handleEditAppointment(e) {
    const id = e.target.getAttribute('data-id');

    try {
        const appointment = await api.get(`/api/appointments/${id}`);

        // Create modal for editing
        const modal = document.createElement('div');
        modal.className = 'modal';
        modal.innerHTML = `
            <div class="modal-content">
                <h2>Edit Appointment</h2>
                <form id="editAppointmentForm">
                    <div class="form-group">
                        <label for="editTitle">Title</label>
                        <input type="text" id="editTitle" value="${appointment.title}" required>
                    </div>
                    <div class="form-group">
                        <label for="editDoctorName">Doctor Name</label>
                        <input type="text" id="editDoctorName" value="${appointment.doctorName}" required>
                    </div>
                    <div class="form-group">
                        <label for="editPatientName">Patient Name</label>
                        <input type="text" id="editPatientName" value="${appointment.patientName}" required>
                    </div>
                    <div class="form-group">
                        <label for="editStartTime">Start Time</label>
                        <input type="datetime-local" id="editStartTime" value="${formatDateTimeForInput(appointment.startTime)}" required>
                    </div>
                    <div class="form-group">
                        <label for="editEndTime">End Time</label>
                        <input type="datetime-local" id="editEndTime" value="${formatDateTimeForInput(appointment.endTime)}" required>
                    </div>
                    <div class="modal-actions">
                        <button type="button" class="btn cancel-btn">Cancel</button>
                        <button type="submit" class="btn">Save Changes</button>
                    </div>
                </form>
            </div>
        `;

        document.body.appendChild(modal);
        modal.style.display = 'flex';

        // Add event listeners
        document.getElementById('editAppointmentForm').addEventListener('submit', async (e) => {
            e.preventDefault();

            const title = document.getElementById('editTitle').value;
            const doctorName = document.getElementById('editDoctorName').value;
            const patientName = document.getElementById('editPatientName').value;
            const startTime = document.getElementById('editStartTime').value;
            const endTime = document.getElementById('editEndTime').value;

            try {
                await api.put(`/api/appointments/${id}`, {
                    title,
                    doctorName,
                    patientName,
                    startTime: new Date(startTime).toISOString(),
                    endTime: new Date(endTime).toISOString()
                });

                modal.remove();
                loadAppointments();
            } catch (error) {
                console.error('Error updating appointment:', error);
                alert('Failed to update appointment');
            }
        });

        document.querySelector('.cancel-btn').addEventListener('click', () => {
            modal.remove();
        });
    } catch (error) {
        console.error('Error fetching appointment:', error);
        alert('Failed to fetch appointment details');
    }
}

async function handleDeleteAppointment(e) {
    if (!confirm('Are you sure you want to delete this appointment?')) {
        return;
    }

    const id = e.target.getAttribute('data-id');

    try {
        await api.delete(`/api/appointments/${id}`);
        loadAppointments();
    } catch (error) {
        console.error('Error deleting appointment:', error);
        alert('Failed to delete appointment');
    }
}

function formatDateTimeForInput(dateTimeString) {
    const date = new Date(dateTimeString);
    const offset = date.getTimezoneOffset() * 60000;
    const localISOTime = new Date(date.getTime() - offset).toISOString().slice(0, 16);
    return localISOTime;
}