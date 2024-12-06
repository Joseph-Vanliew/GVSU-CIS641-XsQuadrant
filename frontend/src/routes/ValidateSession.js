import axios from "axios";

export const validateSession = async () => {
    try {
        const response = await axios.get(`/api/validate`, {
            withCredentials: true, // Including cookies in the request
        });
        console.log("Response from /api/validate:", response.data);
        return response.data.user;
    } catch (error) {
        console.error("Session validation failed:", error.response?.data || error.message);
        return null;
    }
};