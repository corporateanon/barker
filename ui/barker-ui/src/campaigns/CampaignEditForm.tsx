import { Button, Grid, FormControlLabel } from '@material-ui/core';
import { Bot, Campaign } from 'barker-api';
import { Field, Form, Formik } from 'formik';
import { TextField, CheckboxWithLabel } from 'formik-material-ui';
import React, { FC } from 'react';

export interface CampaignEditFormProps {
    campaign: Campaign;
    onSubmit: (bot: Bot) => Promise<void>;
}

const CampaignEditForm: FC<CampaignEditFormProps> = ({
    campaign,
    onSubmit,
}) => {
    return (
        <Grid container>
            <Grid item>
                <Formik initialValues={campaign} onSubmit={onSubmit}>
                    <Form>
                        <Grid container spacing={2}>
                            <Grid item xs={12}>
                                <Field
                                    disabled
                                    component={TextField}
                                    label="ID"
                                    name="ID"
                                    variant="outlined"
                                />
                            </Grid>
                            <Grid item xs={12}>
                                <Field
                                    component={TextField}
                                    label="Title"
                                    name="Title"
                                    variant="outlined"
                                />
                            </Grid>
                            <Grid item xs={12}>
                                <Field
                                    component={TextField}
                                    label="Message"
                                    name="Message"
                                    variant="outlined"
                                />
                            </Grid>
                            <Grid item xs={12}>
                                <Field
                                    component={CheckboxWithLabel}
                                    type="checkbox"
                                    Label={{ label: 'Active' }}
                                    name="Active"
                                />
                            </Grid>
                            <Grid item xs={12}>
                                <Button
                                    type="submit"
                                    color="primary"
                                    variant="contained"
                                >
                                    Save
                                </Button>
                            </Grid>
                        </Grid>
                    </Form>
                </Formik>
            </Grid>
        </Grid>
    );
};
export default CampaignEditForm;
