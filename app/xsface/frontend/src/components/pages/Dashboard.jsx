import React from "react";
import Chat from "../chat.jsx";
import Room from "../room.jsx";
import Stream from "../stream.jsx";
import { Box, Button, Typography, Stack } from "@mui/material";
import { styled } from "@mui/material/styles";
import AppTheme from "../../shared-theme/AppTheme.jsx";
import CssBaseline from "@mui/material/CssBaseline";
import ColorModeSelect from "../../shared-theme/ColorModeSelect.jsx";
import { useNavigate } from "react-router-dom";
import axios from "axios";

const DashboardContainer = styled(Stack)(({ theme }) => ({
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



export default function Dashboard({
                                      roomWebsocketAddr,
                                      chatWebsocketAddr,
                                      viewerWebsocketAddr,
                                      streamWebsocketAddr,
                                      streamLink,
                                      noStream,
                                  }) {
    const navigate = useNavigate();

    const handleLogout = async () => {
        try {
            // Revoking the cookie
            await axios.post(`${import.meta.env.VITE_API_URL}/api/logout`, {}, { withCredentials: true });

            // Remove token from local storage
            localStorage.removeItem("Authorization");

            // Redirect to the landing page
            navigate("/");
        } catch (error) {
            console.error("Logout failed:", error);
        }
    };

    return (
        <AppTheme>
            <CssBaseline enableColorScheme />
            <DashboardContainer direction="column" justifyContent="space-between">
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
                    {/* Header */}
                    <Box
                        sx={{
                            display: "flex",
                            justifyContent: "space-between",
                            alignItems: "center",
                            width: "100%",
                            maxWidth: "1200px",
                            marginBottom: 4,
                        }}
                    >
                        <Typography
                            variant="h4"
                            component="h1"
                            sx={{
                                fontSize: 'clamp(2rem, 10vw, 2.15rem)',
                            }}
                        >
                            Meeting
                        </Typography>
                        <Button
                            variant="outlined"
                            color="secondary"
                            onClick={handleLogout}
                        >
                            Logout
                        </Button>
                    </Box>

                    {/* Content Layout */}
                    <Stack
                        direction={{ xs: "column", md: "row" }}
                        spacing={4}
                        sx={{
                            width: "100%",
                            maxWidth: "1200px",
                        }}
                    >
                        {/* Chat Section */}
                        <Box
                            sx={{
                                flex: 1,
                                display: "flex",
                                flexDirection: "column",
                                overflow: "auto",
                                border: "1px solid #ccc",
                                borderRadius: 1,
                                padding: 2,
                                backgroundColor: "background.paper", // Matches theme
                            }}
                        >
                            <Typography variant="h6" component="h2" sx={{ mb: 2 }}>
                                Chat
                            </Typography>
                            <Chat ChatWebsocketAddr={chatWebsocketAddr} />
                        </Box>

                        {/* Room Section */}
                        <Box
                            sx={{
                                flex: 2,
                                display: "flex",
                                flexDirection: "column",
                                overflow: "auto",
                                border: "1px solid #ccc",
                                borderRadius: 1,
                                padding: 2,
                                backgroundColor: "background.paper", // Matches theme
                            }}
                        >
                            <Typography variant="h6" component="h2" sx={{ mb: 2 }}>
                                Room
                            </Typography>
                            <Room
                                streamLink={streamLink}
                                roomWebsocketAddr={roomWebsocketAddr}
                                chatWebsocketAddr={chatWebsocketAddr}
                                viewerWebsocketAddr={viewerWebsocketAddr}
                            />
                        </Box>

                        {/* Stream Section */}
                        <Box
                            sx={{
                                flex: 1,
                                display: "flex",
                                flexDirection: "column",
                                overflow: "auto",
                                border: "1px solid #ccc",
                                borderRadius: 1,
                                padding: 2,
                                backgroundColor: "background.paper", // Matches theme
                            }}
                        >
                            <Typography variant="h6" component="h2" sx={{ mb: 2 }}>
                                Stream
                            </Typography>
                            <Stream
                                noStream={noStream}
                                streamWebsocketAddr={streamWebsocketAddr}
                                chatWebsocketAddr={chatWebsocketAddr}
                                viewerWebsocketAddr={viewerWebsocketAddr}
                            />
                        </Box>
                    </Stack>
                </Box>
            </DashboardContainer>
        </AppTheme>
    );
}