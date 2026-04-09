import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import Authentication from '../../pages/AuthenticationPage/Authentication';

const AppRouter = () => {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/auth" element={<Authentication />} />
                <Route path="/*" element={<Navigate to="/auth" replace />} />
            </Routes>
        </BrowserRouter>
    );
};

export default AppRouter;