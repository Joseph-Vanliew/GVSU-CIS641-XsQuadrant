import * as React from 'react';
import { Box, Button, Typography, Stack } from '@mui/material';
import { styled } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';
import AppTheme from '../../shared-theme/AppTheme.jsx';
import ColorModeSelect from '../../shared-theme/ColorModeSelect.jsx';
import { useNavigate } from 'react-router-dom';
import {validateSession} from "../../routes/ValidateSession.js";
import {useEffect} from "react";

const LandingContainer = styled(Stack)(({ theme }) => ({
    height: '100vh',
    justifyContent: 'center',
    alignItems: 'center',
    padding: theme.spacing(4),
    backgroundImage:
        'radial-gradient(circle, hsl(210, 100%, 97%), hsl(0, 0%, 100%))',
    backgroundRepeat: 'no-repeat',
    backgroundSize: 'cover',
    ...theme.applyStyles('dark', {
        backgroundImage:
            'radial-gradient(circle, hsla(210, 100%, 16%, 0.5), hsl(220, 30%, 5%))',
    }),
}));

export default function LandingPage(props) {
    const navigate = useNavigate();

    useEffect(() => {
        const checkSession = async () => {
            const user = await validateSession();
            if (user) {
                navigate('/home'); // Redirect if session is valid
            }
        };

        checkSession();
    }, [navigate]);

    return (
        <AppTheme {...props}>
            <CssBaseline enableColorScheme />
            <ColorModeSelect sx={{ position: 'fixed', top: '1rem', right: '1rem' }} />
            <LandingContainer>
                <Typography
                    variant="h3"
                    component="h1"
                    sx={{ fontSize: 'clamp(2.5rem, 5vw, 3.5rem)', textAlign: 'center' }}
                >
                    XsQuadrant
                </Typography>
                <Typography
                    variant="subtitle1"
                    sx={{ textAlign: 'center', marginBottom: 4, maxWidth: '600px' }}
                >
                    A powerful platform for real-time collaboration.
                </Typography>
                <Box sx={{ display: 'flex', gap: 2 }}>
                    <Button
                        variant="contained"
                        size="large"
                        onClick={() => navigate('/signin')}
                        sx={{ minWidth: '150px' }}
                    >
                        Sign In
                    </Button>
                    <Button
                        variant="outlined"
                        size="large"
                        onClick={() => navigate('/signup')}
                        sx={{ minWidth: '150px' }}
                    >
                        Sign Up
                    </Button>
                </Box>
            </LandingContainer>
        </AppTheme>
    );
}