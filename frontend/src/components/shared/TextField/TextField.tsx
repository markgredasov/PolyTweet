import React, { FC, InputHTMLAttributes } from 'react';
import styles from './TextField.module.scss';

interface TextFieldProps extends InputHTMLAttributes<HTMLInputElement> {
    label?: string;
    error?: string;
    touched?: boolean;
}

const TextField: FC<TextFieldProps> = ({ label, error, touched, ...props }) => {
    return (
        <div className={styles.container}>
            {label && <label className={styles.label}>{label}</label>}
            <input className={`${styles.input} ${touched && error ? styles.error : ''}`} {...props} />
            {touched && error && <span className={styles.errorMessage}>{error}</span>}
        </div>
    );
};

export default TextField;