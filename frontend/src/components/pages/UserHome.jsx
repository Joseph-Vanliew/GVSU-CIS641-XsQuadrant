import React from "react";
import { Box, Button, Typography, Stack } from "@mui/material";
import { styled } from "@mui/material/styles";
import { useNavigate } from "react-router-dom";
import AppTheme from "../../shared-theme/AppTheme";
import ColorModeSelect from "../../shared-theme/ColorModeSelect";

const UserHomeContainer = styled(Stack)(({ theme }) => ({
    height: 'calc((1 - var(--template-frame-height, 0)) * 100dvh)',
    minHeight: '100%',
    padding: theme.spacing(2),
    [theme.breakpoints.up('sm')]: {
        padding: theme.spacing(4),
    },
    '&::before': {
        content: '""',
        display: 'block',
        position: 'absolute',
        zIndex: -1,
        inset: 0,
        backgroundImage:
            'radial-gradient(ellipse at 50% 50%, hsl(210, 100%, 97%), hsl(0, 0%, 100%))',
        backgroundRepeat: 'no-repeat',
        ...theme.applyStyles('dark', {
            backgroundImage:
                'radial-gradient(at 50% 50%, hsla(210, 100%, 16%, 0.5), hsl(220, 30%, 5%))',
        }),
    },
}));

const UserHome = () => {
    const navigate = useNavigate();

    const handleCreateMeeting = () => {
        navigate("/room"); // Adjust the route based on your actual configuration
    };

    const handleScheduleMeeting = () => {
        console.log("Schedule Meeting clicked");
    };

    const handleViewMeetings = () => {
        console.log("Calendar clicked");
    };

    return (
        <AppTheme>
            <UserHomeContainer direction="column" justifyContent="space-between">
                <ColorModeSelect sx={{ position: 'fixed', top: '1rem', right: '1rem' }} />
                <Box
                    sx={{
                        display: "flex",
                        flexDirection: "column",
                        alignItems: "center",
                        gap: 4,
                        width: "100%",
                    }}
                >
                    <Typography variant="h4" gutterBottom>
                        Home
                    </Typography>
                    <Button
                        variant="contained"
                        color="primary"
                        onClick={handleCreateMeeting}
                        sx={{ mb: 2 }}
                    >
                        Create Meeting
                    </Button>
                    <Button
                        variant="contained"
                        color="secondary"
                        onClick={handleScheduleMeeting}
                        sx={{ mb: 2 }}
                    >
                        Schedule Meeting
                    </Button>
                    <Button variant="contained" onClick={handleViewMeetings}>
                        Calendar
                    </Button>
                </Box>
            </UserHomeContainer>
        </AppTheme>
    );
};

export default UserHome;